package main

import (
	"fmt"
	"math/rand"
)

type Cell struct {
	symbol string
	status string
}

const ROWS, COLS int = 10, 15
var cave = make([][]Cell, ROWS)
var posX, posY int = 1, 1

var wall string = "🧱"
var dark string = "⬛"
var discovered string = "  "
var player string = "🙂"
var treasure string = "💎"

func generateCave() {
	for y:=0; y<ROWS; y++ {
		cave[y] = make([]Cell, COLS)
		for x:=0; x<COLS; x++ {
			if y == 0 || y == ROWS - 1 || x == 0 || x == COLS - 1 {
				cave[y][x].symbol = wall
				cave[y][x].status = "safe"
			} else {
				cave[y][x].symbol = dark
				cave[y][x].status = "safe"
				
				if rand.Intn(10) == 0 {
					cave[y][x].status = "death"
				}
			}
		}
	}

	cave[posY][posX].symbol = player
	cave[posY][posX].status = "safe"

	cave[ROWS - 2][COLS - 2].symbol = treasure
	cave[ROWS - 2][COLS - 2].status = "safe"
}

func displayCave() {
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
			return cave[posY - 1][posX].symbol != wall
		case "a":
			return cave[posY][posX - 1].symbol != wall
		case "s":
			return cave[posY + 1][posX].symbol != wall
		case "d":
			return cave[posY][posX + 1].symbol != wall
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
			if cave[newY][newX].status == "death" {
				fmt.Println("You just found the death cell 💀")
				break
			} else if cave[newY][newX].symbol == "💎" {
				fmt.Println("Congratulations! You have won the game 👏")
				break
			}

			cave[posY][posX].symbol = discovered
			posY, posX = newY, newX
			cave[posY][posX].symbol = player
		}
	}
}