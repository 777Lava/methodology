package games

import (
	"fmt"
	"math"
	"math/rand"
)

func Game1(userName string) {
	println("Find the smallest common multiple of given numbers.")
	for {
		a, b, c := rand.Intn(100), rand.Intn(100), rand.Intn(100)

		fmt.Printf("Question: %d %d %d\n", a, b, c)
		fmt.Print("Your answer: ")
		var answerUser, answer int
		_, err := fmt.Scan(&answerUser)
		if err != nil {
			println("Invalid input. let's try again")
		}
		answer = (lcm(lcm(a, b), c))
		if answerUser == answer {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("'%d' is wrong answer ;(. Correct answer was '%d'.\nLet's try again, %s!\n", answerUser, answer, userName)
		}
	}
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b))) / gcc(a, b)
}
func gcc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
