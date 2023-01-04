package main

type alienAdapter struct {
	alien alienMachine
}

func (a *alienAdapter) insertInSquarePort() {
	a.alien.insertInSquarePort()
}
