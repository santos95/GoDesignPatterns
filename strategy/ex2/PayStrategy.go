package main

type PayStrategy interface {
	pay(paymentAmount int) bool
	collectPaymentDetails()
}
