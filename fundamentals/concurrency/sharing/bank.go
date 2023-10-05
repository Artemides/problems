package sharing

import "fmt"

var (
	deposits = make(chan int)
	balances = make(chan int)
)

func deposit(amount int) { deposits <- amount }
func balance() int       { return <-balances }

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func BankMain() {
	go teller()
	deposit(10)
	deposit(15)
	go deposit(30)
	go deposit(40)
	fmt.Println(balance())

}
