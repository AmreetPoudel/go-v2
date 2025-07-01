package main

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func main() {
	fmt.Println("welcome to dice game simulator")
	fmt.Println("we will be using two dices")
	for {
		HomeScreen()
		var choice int
		fmt.Scan(&choice)
		if choice != 1 && choice != 2 {
			fmt.Println("invalid choice ")
			fmt.Println("exiting from game")
			break
		}
		if choice == 2 {
			fmt.Println("Thank you for playing")
			break
		}
		RollDice()

	}
}
func HomeScreen() {
	fmt.Println("enter your choice")
	fmt.Println("1. want to roll a dice")
	fmt.Println("2. exit from game.")
	fmt.Print("enter your choice:::  ")

}

func RollDice() {
	n1 := rand.Intn(6) + 1
	n2 := rand.Intn(6) + 1
	fmt.Printf("dice number are %d and %d \nso sum is : %d\n", n1, n2, n1+n2)

}
