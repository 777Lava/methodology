package both

import "fmt"

func Greetings() string{ 

	var userName string
	println("Welcome to the Brain Games!")
	print("May I have your name? ")
	_, err := fmt.Scan(&userName)
	if err != nil {
		println("Invalid input. let's try again")
	}
	fmt.Printf("Hello, %s!\n", userName)
	return userName
}