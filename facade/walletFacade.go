package main

import (
	"fmt"
)

type walletFacade struct {
	account      *account
	securityCode *securityCode
	wallet       *wallet
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountID string, code int) *walletFacade {

	fmt.Println("Starting create account")

	walletFacade := &walletFacade{
		account:      newAccount(accountID),
		securityCode: createSecurityXode(code),
		wallet:       newWallet(),
		notification: &notification{},
		ledger:       &ledger{},
	}

	fmt.Println("Account Created")
	return walletFacade
}

func (w *walletFacade) addMoneyToWaller(accountID string, code int, amount int) error {

	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountID)

	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)

	if err != nil {
		return err
	}

	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountID, "credit", amount)

	return nil

}

func (w *walletFacade) deductMoneyFromWallet(accountID string, code int, amount int) error {

	fmt.Println("Starting debit money from wallet")
	err := w.account.checkAccount(accountID)

	if err != nil {
		return err
	}

	err = w.securityCode.checkCode(code)

	if err != nil {
		return err
	}

	w.wallet.debitBalance(amount)
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountID, "debit", amount)
	return nil

}

func (w *walletFacade) balance() int {
	return w.wallet.getBalance()
}
