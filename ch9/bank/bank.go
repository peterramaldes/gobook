// The bank packate implements bank with only one account
package bank

var balance int

func Deposit(amount int) { balance = balance + amount }
func Balance() int       { return balance }
