package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("------------------------------------------------------------------------------------------")
	walletFacade := newWalletFacade("santos", 1234)
	fmt.Println("------------------------------------------------------------------------------------------")
	err := walletFacade.addMoneyToWaller("santos", 1234, 500)

	if err != nil {
		log.Fatalf("Error %s\n", err)
	}

	fmt.Println("------------------------------------------------------------------------------------------")
	err = walletFacade.deductMoneyFromWallet("santos", 1234, 300)

	if err != nil {
		log.Fatalf("Error %s\n", err)
	}

	fmt.Println("------------------------------------------------------------------------------------------")
	fmt.Printf("Balance: %d\n", walletFacade.balance())
	fmt.Println("------------------------------------------------------------------------------------------")

}
