package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	salario := "1000"
	salarioConvertido, _ := strconv.Atoi(salario) //Atoi = ASCII to Int

	fmt.Println(salario, reflect.TypeOf(salario))
	fmt.Println(salarioConvertido, reflect.TypeOf(salarioConvertido))

	inteiro := 25
	vinteCinco := strconv.Itoa(inteiro) //Itoa = Int to ASCII

	fmt.Println(inteiro, reflect.TypeOf(inteiro))
	fmt.Println(vinteCinco, reflect.TypeOf(vinteCinco))
}
