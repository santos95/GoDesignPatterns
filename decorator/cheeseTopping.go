package main

type SheeseTopping struct {
	pizza ipizza
}

func (c *SheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}
