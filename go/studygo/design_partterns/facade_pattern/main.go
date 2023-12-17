package main

import (
	"fmt"
	"log"
	fa "studygo/design_partterns/facade_pattern/src"
)

func main() {
	fmt.Println()
	walletFacade := fa.NewWalletFacade("abc", 1234)
	fmt.Println()
	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}
	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("error: %s \n", err.Error())
	}
}
