package main

import (
	"fmt"
	"time"
)

func main() {

	data := "23/Mar/2025 - 16:13:00"
	fmt.Println(data)

	agora := time.Now()
	fmt.Println(agora)

	//format date golang
	agoraFormat := agora.Format("02/01/2006 15:04:05")
	fmt.Println(agoraFormat)

	agoraFormat2 := agora.Format("02/Jan/2006 - 15:04:05")
	fmt.Println(agoraFormat2)

	//dia (com zero) - 02
	//dia (sem zero) - 2
	//mês (com zero) - 01
	//mês (sem zero) - 1
	//ano (com 4 dígitos) - 2006
	//hora (24h) - 15
	//hora (12h) - 03
	//minutos - 04
	//segundos - 05
	//AM/PM - PM
	//fuso horário - MST
	//dia semana inteiro - Monday
	//dia semana abreviado - Mon
	//mês inteiro - January
	//mês abreviado - Jan

	dia := time.Now().Format("02")
	mes := time.Now().Format("01")
	ano := time.Now().Format("2006")
	fmt.Println(dia, mes, ano)

	hora := time.Now().Format("15")
	minuto := time.Now().Format("04")
	segundo := time.Now().Format("05")
	fmt.Println(hora, minuto, segundo)

	//dia da semana
	diaSemana := time.Now().Format("Monday")
	fmt.Println(diaSemana)

	//parse date golang
	dataParse, _ := time.Parse("02/Jan/2006 - 15:04:05", data)
	fmt.Println(dataParse)

}
