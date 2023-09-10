package game

import (
	"fmt"
	"os"
	"os/exec"
	"math/rand"
)

func Clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func DisplayCave() {
	Clear()
	
	for y:=0; y<ROWS; y++ {
		for x:=0; x<COLS; x++ {
			fmt.Print(cave[y][x].symbol)
		}
		fmt.Println()
	}
}

func PlayerDied() {
	cave[posY][posX].symbol = DIED
	DisplayCave()
	fmt.Println("You just found the death cell 💀")
	fmt.Println()
}

func PlayerWon() {
	cave[posY][posX].symbol = WON
	DisplayCave()
	fmt.Println("Congratulations! You have won the game 👏")
	fmt.Println()
}

func PlayAgain() {
	var choice string
	for choice != "Yes" && choice != "No" {
		fmt.Print("Play again? [Yes|No]: ")
		fmt.Scan(&choice)
	}

	switch choice {
		case "Yes":
			posX, posY = 1, 1
			StartGame()
		case "No":
			MainMenu()
	}
}

func ResetCells() {
	for y:=1; y<ROWS-1; y++ {
		for x:=1; x<COLS-1; x++ {
			if cave[y][x].symbol == PLAYER || cave[y][x].symbol == TREASURE || cave[y][x].symbol == DISCOVERED {
				continue
			}
			cave[y][x].symbol = DARK
		}
	}
}

func ShuffleDeathCells() {
	for y:=1; y<ROWS-1; y++ {
		for x:=1; x<COLS-1; x++ {
			if cave[y][x].symbol == PLAYER || cave[y][x].symbol == TREASURE || cave[y][x].symbol == DISCOVERED {
				continue
			}
			if rand.Intn(8) == 0 {
				cave[y][x].status = "danger"
			}
		}
	}
}