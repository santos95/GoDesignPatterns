package main

import "fmt"

type ledger struct{}

func (l *ledger) makeEntry(accountId, txtType string, amount int) {
	fmt.Printf("Make ledger entry for accountId %s with txtnType %s for amount %d\n", accountId, txtType, amount)
	return
}
