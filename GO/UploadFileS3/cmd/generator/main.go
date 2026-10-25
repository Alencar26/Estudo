package main

import (
	"fmt"
	"os"
)

func main() {

	i := 0
	for {
		file, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		file.WriteString("Olá Mundo!")
		file.Close()
		i++
	}
}
