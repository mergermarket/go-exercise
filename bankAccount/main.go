package bankAccount

type bankAccount struct {

}


func main() {

}


type Account interface {

	// Withdraw deducts amount from account returning the new balance
	Withdraw(amount int) (newBalance int)

	// Deposit adds amount to the account returning the new balance
	Deposit(amount int) (newBalance int)
}


func
