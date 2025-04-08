package dompet

type Wallet struct {
	owner   string
	balance int
}

func RegisterWallet(username string) Wallet {
	return Wallet{owner: username, balance: 0}
}

func (wallet *Wallet) GetInfo() int {
	return wallet.balance
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.balance = amount
}
