package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeshort := nikeFactory.makeShort()

	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()

	fmt.Println("Shoes")
	printShoeDetails(adidasShoe)
	printShoeDetails(nikeShoe)

	fmt.Println("Shorts")
	printShortDetails(nikeshort)
	printShortDetails(adidasShort)

}

func printShoeDetails(s iShoe) {

	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()

}

func printShortDetails(s iShort) {

	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()

}
