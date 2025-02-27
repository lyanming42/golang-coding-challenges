//  Design a simple bank account system with the ability to:
// - Deposit money
// - Withdraw money
// - Get the current balance

package main

import (
	"errors"
	"fmt"
)

type BankAccount struct {
	balance int
}

func (account BankAccount) deposit(amount int) BankAccount {
	account.balance += amount
	return account
}

func (account BankAccount) withdraw(amount int) (BankAccount, error) {
	fmt.Println(account.balance)
	fmt.Println(amount)
	if account.balance < amount {
		return account, errors.New("No enough balance for account")
	}
	account.balance -= amount
	return account, nil
}

func (account BankAccount) getBalance() (BankAccount, int) {
	return account, account.balance
}

func main() {
	var account BankAccount
	var value int
	account = account.deposit(10)
	account, err := account.withdraw(5)
	if err != nil {
		fmt.Println("No enough balance")
	}
	account, value = account.getBalance()
	fmt.Println(value)
	account, err = account.withdraw(10)
	if err != nil {
		fmt.Println("No enough balance")
	}
	fmt.Println(account)
	fmt.Println("Hello World")
}
