package main

import (
	"database/sql"
	"net"

	"github.com/Alencar26/Estudo/GO/GRPC/internal/database"
	"github.com/Alencar26/Estudo/GO/GRPC/internal/pb"
	"github.com/Alencar26/Estudo/GO/GRPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	//Criando servidor GRPC
	grpcServer := grpc.NewServer()

	//Vamos usar um client GRPC chamado Evans
	//https://github.com/ktr0731/evans
	//Para que ele funcione, precisamos usar Reflection. (Reflection é quando ele consegue ler e processsar sua própria informação)
	reflection.Register(grpcServer)

	//para rodar o Evans execute no terminal: evans -r repl
	//Comando Evans:
	// packge <seu_pacote> - para definir o pacote que será usando
	// service <seu-servico> - para definir qual serviço usar
	// call <nome_da_sua_chamada> - para definir qual chamada executar

	//Registrando meu serviço no servidor GRPC
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)

	//É preciso abrir uma conexão TCP para conseguir se comunicar com GRPC
	//Abrindo conexão TCP
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
