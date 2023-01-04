package main

type TomatoTopping struct {
	pizza ipizza
}

func (c *TomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}
