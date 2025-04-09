package dompet

type Wallet struct {
	Owner   string
	Balance int
}

func RegisterWallet(username string) Wallet {
	return Wallet{Owner: username, Balance: 0}
}

func (wallet *Wallet) GetInfo() int {
	return wallet.Balance
}

func (wallet *Wallet) Deposit(amount int) {
	wallet.Balance = amount
}
