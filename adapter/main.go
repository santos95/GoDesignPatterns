package main

func main() {
	client := &client{}
	mac := &mac{}
	client.insertInSquarePort(mac)

	windowsMachine := &windows{}
	windowsMachineAdapter := &windowsAdapter{
		windowsMachine: windowsMachine,
	}
	client.insertInSquarePort(windowsMachineAdapter)

	alienMachine := &alienMachine{}
	alienMachineAdapter := &alienAdapter{
		alien: alienMachine,
	}
	client.insertInSquarePort(alienMachineAdapter)
}