package main

import (
	"fmt"
	"os"
	"os/exec"
	"math/rand"
)

type Cell struct {
	symbol string
	status string
}

const ROWS, COLS int = 10, 15
var cave = make([][]Cell, ROWS)
var posX, posY int = 1, 1

const (
	WALL       string = "🧱"
	DARK       string = "⬛"
	DISCOVERED string = "  "
	PLAYER     string = "😨"
	DIED       string = "😵"
	WON        string = "😁"
	TREASURE   string = "💎"
)

func clear() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func generateCave() {
	for y:=0; y<ROWS; y++ {
		cave[y] = make([]Cell, COLS)
		for x:=0; x<COLS; x++ {
			if y == 0 || y == ROWS - 1 || x == 0 || x == COLS - 1 {
				cave[y][x].symbol = WALL
				cave[y][x].status = "safe"
			} else {
				cave[y][x].symbol = DARK
				cave[y][x].status = "safe"
				
				if rand.Intn(10) == 0 {
					cave[y][x].status = "death"
				}
			}
		}
	}

	cave[posY][posX].symbol = PLAYER
	cave[posY][posX].status = "safe"

	cave[ROWS - 2][COLS - 2].symbol = TREASURE
	cave[ROWS - 2][COLS - 2].status = "safe"
}

func displayCave() {
	clear()
	
	for y:=0; y<ROWS; y++ {
		for x:=0; x<COLS; x++ {
			fmt.Print(cave[y][x].symbol)
		}
		fmt.Println()
	}
}

func isNotWall(move string) bool {
	switch move {
		case "w":
			return cave[posY - 1][posX].symbol != WALL
		case "a":
			return cave[posY][posX - 1].symbol != WALL
		case "s":
			return cave[posY + 1][posX].symbol != WALL
		case "d":
			return cave[posY][posX + 1].symbol != WALL
		default:
			return false
	}
}

func getNewPosition(move string) (int, int) {
	switch move {
		case "w":
			return posY - 1, posX
		case "a":
			return posY, posX - 1
		case "s":
			return posY + 1, posX
		case "d":
			return posY, posX + 1
		default:
			return posY, posX
	}
}

func main() {
	var move string
	generateCave()

	for {
		displayCave()
		fmt.Print("Input move [w|a|s|d]: ")
		fmt.Scan(&move); fmt.Scanln()

		if isNotWall(move) {
			newY, newX := getNewPosition(move)

			cave[posY][posX].symbol = DISCOVERED
			posY, posX = newY, newX

			if cave[newY][newX].status == "death" {
				cave[posY][posX].symbol = DIED
				displayCave()
				fmt.Println("You just found the death cell 💀")
				break
			} else if cave[newY][newX].symbol == TREASURE {
				cave[posY][posX].symbol = WON
				displayCave()
				fmt.Println("Congratulations! You have won the game 👏")
				break
			}

			cave[posY][posX].symbol = PLAYER
		}
	}
}