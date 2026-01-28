package main

import (
	"fmt"
)

func main() {

	adidasFactory, err := GetSportFactory("adidas")
	checkErr(err)

	nikeFactor, err := GetSportFactory("nike")
	checkErr(err)

	adidasShoe := adidasFactory.makeShoe()
	adidasShirt := adidasFactory.makeShirt()

	printShoeDatails(adidasShoe)
	printShirtDatails(adidasShirt)

	nikeShoe := nikeFactor.makeShoe()
	nikeShirt := nikeFactor.makeShirt()

	printShoeDatails(nikeShoe)
	printShirtDatails(nikeShirt)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func printShoeDatails(s IShoe) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}

func printShirtDatails(s IShirt) {
	fmt.Printf("Logo: %s\n", s.getLogo())
	fmt.Printf("Size: %d\n", s.getSize())
}
