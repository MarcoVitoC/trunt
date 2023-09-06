package main

import "fmt"

const ROWS, COLS int = 10, 15
var cave = make([][]string, ROWS, COLS)
var posX, posY int = 1, 1

var wall string = "🧱"
var dark string = "⬛"
var revealed string = "  "
var player string = "🙂"
var treasure string = "💎" 

func generateCave() {
	for y:=0; y<ROWS; y++ {
		cave[y] = make([]string, COLS)
		for x:=0; x<COLS; x++ {
			if y == 0 || y == ROWS - 1 || x == 0 || x == COLS - 1 {
				cave[y][x] = wall
			} else {
				cave[y][x] = dark
			}
		}
	}

	cave[posY][posX] = player
	cave[ROWS - 2][COLS - 2] = treasure
}

func displayCave() {
	for y:=0; y<ROWS; y++ {
		for x:=0; x<COLS; x++ {
			fmt.Print(cave[y][x])
		}
		fmt.Println()
	}
}

func main() {
	var move string
	generateCave()

	for {
		displayCave()
		fmt.Print("Input move [w|a|s|d]: ")
		fmt.Scan(&move); fmt.Scanln()

		switch move {
			case "w":
				if cave[posY - 1][posX] != wall {
					cave[posY][posX] = revealed
					posY--
					cave[posY][posX] = player
				}
			case "a":
				if cave[posY][posX - 1] != wall {
					cave[posY][posX] = revealed
					posX--
					cave[posY][posX] = player
				}
			case "s":
				if cave[posY + 1][posX] != wall {
					cave[posY][posX] = revealed
					posY++
					cave[posY][posX] = player
				}
			case "d":
				if cave[posY][posX + 1] != wall {
					cave[posY][posX] = revealed
					posX++
					cave[posY][posX] = player
				}
		}
	}
}