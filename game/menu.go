package game

import (
	"fmt"
	"os"
)

func howToPlay() {
	Clear()
	fmt.Println("How to play")
	fmt.Println("===========")
	fmt.Println("1. Your goal is to reach the treasure (💎).")
	fmt.Println("2. Move the player (😨) by pressing 'w' to move up, 'a' to move left, 's' to move down, 'd' to move right.")
	fmt.Println("3. Beware of death cells among all black cells.")
	fmt.Println("4. Death cells are shuffled every 5 seconds.")
	fmt.Println()
	fmt.Print("Press enter to continue...")
	fmt.Scanln()
	MainMenu()
}

func MainMenu() {
	var choice int

	for choice < 1 || choice > 3 {
		Clear()
		fmt.Println("🗺️  Trunt")
		fmt.Println("=========")
		fmt.Println("1. Start game")
		fmt.Println("2. How to play")
		fmt.Println("3. Exit")
		fmt.Print(">> ")
		fmt.Scan(&choice)
		fmt.Scanln()
	}

	switch choice {
		case 1:
			posX, posY = 1, 1
			StartGame()
		case 2:
			howToPlay()
		case 3:
			fmt.Println()
			fmt.Println("Thank you for playing! 🙂")
			os.Exit(0)
	}
}