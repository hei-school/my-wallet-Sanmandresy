package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Wallet struct {
	balance float64
}

func NewWallet() *Wallet {
	return &Wallet{}
}

func (w *Wallet) GetBalance() float64 {
	return w.balance
}

func (w *Wallet) Deposit(amount float64) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount float64) {
	if amount > w.balance {
		fmt.Println("Balance not enough for this withdrawal")
	} else {
		w.balance -= amount
	}
}

func displayMenu(wallet *Wallet) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("What do you want to do ?")
		fmt.Println("1) See balance")
		fmt.Println("2) Deposit")
		fmt.Println("3) Withdraw")
		fmt.Println("4) Never mind")
		fmt.Println("Enter your choice: ")
		text, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(text)

		switch choice {
		case "1":
			fmt.Println("Your balance is: ", wallet.GetBalance())
		case "2":
			fmt.Print("Enter the amount to deposit: ")
			text, _ = reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(text), 64)
			wallet.Deposit(amount)
		case "3":
			fmt.Print("Enter the amount to withdraw: ")
			text, _ = reader.ReadString('\n')
			amount, _ := strconv.ParseFloat(strings.TrimSpace(text), 64)
			wallet.Withdraw(amount)
		case "4":
			fmt.Println("Bye !")
			return
		}
	}
}

func main() {
	wallet := NewWallet()
	displayMenu(wallet)
}
