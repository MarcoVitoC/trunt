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
				cave[y][x].status = ""
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
	cave[ROWS - 2][COLS - 2].symbol = treasure
}

func displayCave() {
	for y:=0; y<ROWS; y++ {
		for x:=0; x<COLS; x++ {
			fmt.Print(cave[y][x].symbol)
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

		if move == "w" {
			if cave[posY - 1][posX].symbol != wall {
				if cave[posY - 1][posX].status == "death" {
					fmt.Println("You just found the death cell and lost. 💀")
					break
				} else if cave[posY - 1][posX].symbol == "💎" {
					fmt.Println("Congratulations! You have won the game. 👏")
					break
				}

				cave[posY][posX].symbol = discovered
				posY--
				cave[posY][posX].symbol = player
			}
		} else if move == "a" {
			if cave[posY][posX - 1].symbol != wall {
				if cave[posY][posX - 1].status == "death" {
					fmt.Println("You just found the death cell and lost. 💀")
					break
				} else if cave[posY][posX - 1].symbol == "💎" {
					fmt.Println("Congratulations! You have won the game. 👏")
					break
				}

				cave[posY][posX].symbol = discovered
				posX--
				cave[posY][posX].symbol = player
			}
		} else if move == "s" {
			if cave[posY + 1][posX].symbol != wall {
				if cave[posY + 1][posX].status == "death" {
					fmt.Println("You just found the death cell and lost. 💀")
					break
				} else if cave[posY + 1][posX].symbol == "💎" {
					fmt.Println("Congratulations! You have won the game. 👏")
					break
				}

				cave[posY][posX].symbol = discovered
				posY++
				cave[posY][posX].symbol = player
			}
		} else if move == "d" {
			if cave[posY][posX + 1].symbol != wall {
				if cave[posY][posX + 1].status == "death" {
					fmt.Println("You just found the death cell and lost. 💀")
					break
				} else if cave[posY][posX + 1].symbol == "💎" {
					fmt.Println("Congratulations! You have won the game. 👏")
					break
				}

				cave[posY][posX].symbol = discovered
				posX++
				cave[posY][posX].symbol = player
			}
		}
	}
}