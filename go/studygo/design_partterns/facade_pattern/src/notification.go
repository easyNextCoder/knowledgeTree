package src

import "fmt"

type Notification struct{}

func (n *Notification) sendWalletCreditNotification() {
	fmt.Println("sending wallet credit notification")
}

func (n *Notification) sendWalletDebitNotification() {
	fmt.Println("Sending wallet debit notification")
}
