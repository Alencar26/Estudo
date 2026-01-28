package main

import (
	"fmt"
)

type ISportFactory interface {
	makeShoe() IShoe
	makeShirt() IShirt
}

func GetSportFactory(brand string) (ISportFactory, error) {
	switch brand {
	case "adidas":
		return &Adidas{}, nil
	case "nike":
		return &Nike{}, nil
	default:
		return nil, fmt.Errorf("invalid brand")
	}
}
