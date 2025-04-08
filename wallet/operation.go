package wallet

import "fmt"

type Wallet struct {
	owner   string
	balance int
}

var myWallet Wallet

func init() {
	myWallet = Wallet{
		owner:   "Ari Susanto",
		balance: 0,
	}
}

func GetInfo() {
	fmt.Println(myWallet)
}
