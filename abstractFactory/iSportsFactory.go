package main

import "fmt"

type iSportFactory interface {
	makeShoe() iShoe
	makeShort() iShort
}

func getSportsFactory(brand string) (iSportFactory, error) {
	if brand == "adidas" {
		return &adidas{}, nil
	} else if brand == "nike" {
		return &nike{}, nil
	}
	return nil, fmt.Errorf("wrong brand type passed")
}
