// The bank package implements bank with only one account, offer concurrency
package bank

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

func teller() {
	var balance int // balance will be confined on goroutine teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // initiate the go routine responsible for monitor
}
