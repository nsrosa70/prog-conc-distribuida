package bank

var balance int

func Deposit(amount int) { balance = balance + amount }

func Balance() int { return balance }

func SetBalance(b int) {balance = 0}



