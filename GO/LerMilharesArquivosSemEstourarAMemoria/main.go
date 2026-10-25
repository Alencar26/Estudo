package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if dir, err := os.Open("/tmp"); err != nil {
		panic(err)
	} else {
		for {
			if files, err := dir.Readdir(1); err != nil {
				if err == io.EOF {
					break
				}
				log.Println("ERROR:", err)
				continue
			} else {
				//colocar lógica aqui para tratativa dos arquivos.
				fmt.Println(files[0].Name())
			}
		}
	}
}
