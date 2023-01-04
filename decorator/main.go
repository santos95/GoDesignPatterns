package main

import "fmt"

func main() {

	pizza := &VeggieMania{}

	pizzaWithCheese := &SheeseTopping{
		pizza: pizza,
	}

	pizzaWithTomatosAndCheese := &TomatoTopping{
		pizza: pizzaWithCheese,
	}

	fmt.Printf("The price of veggeMania with tomato and cheese toppings is %d\n", pizzaWithTomatosAndCheese.getPrice())

	pizza2 := &VeggieMania{}

	pizzaWithCheese2 := &SheeseTopping{
		pizza: pizza2,
	}

	fmt.Printf("The price of veggeMania with cheese toppings is %d\n", pizzaWithCheese2.getPrice())

	fmt.Printf("The price of veggeMania with tomato and cheese toppings is %d\n", pizzaWithTomatosAndCheese.getPrice())

}
