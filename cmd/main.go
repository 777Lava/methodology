package main

import (
	"fmt"
	"methodology/both"
	"methodology/games"
)

func main() {
	userName := both.Greetings()
	var game int
	fmt.Print("choose the game: 1 or 2 ")
	_, err := fmt.Scanf("%d\n", &game)
	if err != nil {
		fmt.Println("invalid input")
	}
	if game == 1 {
		for {
			games.Game1(userName)
		}
	} else if game == 2 {
		games.Game2(userName)
	} else {
		fmt.Println("try again")
	}
}
