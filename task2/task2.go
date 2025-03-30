package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var userName string
	println("Welcome to the Brain Games!")
	print("May I have your name? ")
	_,err := fmt.Scan(&userName)
	if err != nil {
		println("Invalid input. let's try again")
	}
	fmt.Printf("Hello, %s!\n", userName)
	println("What number is missing in the progression?")
	for {
		b1 := rand.Intn(10-1) + 1
		q := rand.Intn(10-1) + 1
		n := rand.Intn(10-3) + 3
		miss := rand.Intn(n-3) + 3
		arr := []int{b1}
		for i := 1; i < n; i++ {
			arr = append(arr, arr[i-1]*q)
		}
		fmt.Printf("Question: ")
		for i := 0; i < len(arr); i++ {
			if i == miss {
				fmt.Print("... ")
			} else {
				fmt.Print(arr[i], " ")
			}
		}
		fmt.Println()
		fmt.Print("Your answer: ")
		var answerUser int
		_, err := fmt.Scan(&answerUser)
		if err != nil {
			println("Invalid input. let's try again")
		}

		if answerUser == arr[miss] {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\nLet's try again, %s!\n", answerUser, arr[miss], userName)
		}
	}
}
