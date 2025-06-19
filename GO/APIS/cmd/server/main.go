package main

import (
	"fmt"

	"github.com/Alencar26/Estudo/GO/APIS/configs"
)

func main() {
	config, _ := configs.LoadConfig(".")
	fmt.Println(config.DBDriver)
}
