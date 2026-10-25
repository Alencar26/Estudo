package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Println("1 - Monitorar\n2 - Logs\n3 - Sair")

	var input int
	fmt.Scanf("%d", &input)

	switch input {
	case 1:
		var dir string
		fmt.Println("Digite o diretórico com arquivo de sites que deseja monitorar:")
		fmt.Scan(&dir)
		for _, site := range lerArquivo(dir, "sites.txt") {
			if site == "" {
				break
			}
			monitorarSite(site)
			time.Sleep(10 * time.Second)
		}
		main()
	case 2:
		var dir string
		fmt.Println("Digite o diretórico com arquivo de sites que deseja monitorar:")
		fmt.Scan(&dir)
		for _, log := range lerArquivo(dir, "log.txt") {
			if log == "" {
				break
			}
			println(log)
		}
		main()
	case 3:
		fmt.Println("Saindo...")
		os.Exit(0) //comando do SO para indicar saída.
	default:
		fmt.Println("Opção inválida")
		os.Exit(-1) //comando do SO para indicar erro.
	}
}

func monitorarSite(site string) {
	fmt.Println("Monitorando... [" + site + "]")
	resp, err := http.Get("http://" + site)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		fmt.Println("Site está no ar!")
		registraLog(site, true)
	} else {
		fmt.Println("Site está fora do ar!")
		registraLog(site, false)
	}
}

func lerArquivo(path, siteFile string) []string {
	file, err := os.Open(path + siteFile)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	r := bufio.NewReader(file)
	sites := []string{}
	for {
		line, _, err := r.ReadLine()
		sites = append(sites, string(line))
		if err == io.EOF {
			break
		}
	}
	return sites
}

func registraLog(site string, status bool) {
	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	w.WriteString("[" + time.Now().Format("02/01/2006 15:04:05") + "] " + site + " - online: " + fmt.Sprint(status) + "\n")
	w.Flush()
}
