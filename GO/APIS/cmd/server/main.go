package main

import (
	"log"
	"net/http"

	"github.com/Alencar26/Estudo/GO/APIS/configs"
	_ "github.com/Alencar26/Estudo/GO/APIS/docs"
	"github.com/Alencar26/Estudo/GO/APIS/internal/entity"
	"github.com/Alencar26/Estudo/GO/APIS/internal/infa/database"
	"github.com/Alencar26/Estudo/GO/APIS/internal/infa/webserver/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
COMANDO NO TERMINAL PARA GERAR OS ARQUIVOS DO SWAGGER A CADA MODIFICAÇÃO:
*
swag init -g /path/do/arquivo/main.go
*
*/

// @title                           Go Project API with Go-Chi
// @version                         1.0
// @description                     Product API with authenticator
// @termsOfService                   http://swagger.io/terms/

// @contact.name                    André Alencar
// @contact.url                     http://github.com/alencar26
// @contact.email                   email.example@email.com

// @license.name                    Alencar26
// @license.url                     http://github.com/alencar26

// @host                            localhost:8000
// @BasePath                        /
// @securityDefinitions.apikey      ApiKeyAuth
// @in                              heather
// @name                            Authorization
func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDB := database.NewProduct(db)
	productHandler := handlers.NewProductHandler(productDB)

	userDB := database.NewUser(db)
	userHandler := handlers.NewUserHandler(userDB)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)                                  //impede o servidor de cair após um exception
	r.Use(middleware.WithValue("jwt", configs.TokenAuth))        //passando dados por contexto através de middlewares
	r.Use(middleware.WithValue("expireIn", configs.JwtExpireIn)) //passando informações por contexto
	//r.Use(LogRequest) //Middlewaree Customizado

	r.Route("/products", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", productHandler.CreateProduct)
		r.Get("/{id}", productHandler.GetProduct)
		r.Put("/{id}", productHandler.UpdateProduct)
		r.Delete("/{id}", productHandler.DeleteProduct)
		r.Get("/", productHandler.GetProducts)

	})

	r.Post("/users", userHandler.CreateClient)
	r.Post("/users/generate_token", userHandler.GetJWT)

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8000/docs/doc.json"),
	))

	http.ListenAndServe(":8000", r)
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
