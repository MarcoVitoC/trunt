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

var wall Cell = Cell{"🧱", ""}
var dark Cell = Cell{"⬛", "safe"}
var isDiscovered Cell = Cell{"  ", ""}
var player Cell = Cell{"🙂", ""}
var treasure Cell = Cell{"💎", ""} 

func generateCave() {
	for y:=0; y<ROWS; y++ {
		cave[y] = make([]Cell, COLS)
		for x:=0; x<COLS; x++ {
			if y == 0 || y == ROWS - 1 || x == 0 || x == COLS - 1 {
				cave[y][x] = wall
			} else {
				cave[y][x] = dark
				
				if rand.Intn(10) == 0 {
					cave[y][x].symbol = "👾"
				}
			}
		}
	}

	cave[posY][posX] = player
	cave[ROWS - 2][COLS - 2] = treasure
}

func displayCave() {
	for y:=0; y<ROWS; y++ {
		for x:=0; x<COLS; x++ {
			fmt.Print(cave[y][x].symbol)
		}
		fmt.Println()
	}
}

func move() {
	var move string

	fmt.Print("Input move [w|a|s|d]: ")
	fmt.Scan(&move); fmt.Scanln()

	switch move {
		case "w":
			if cave[posY - 1][posX] != wall {
				cave[posY][posX] = isDiscovered
				posY--
				cave[posY][posX] = player
			}
		case "a":
			if cave[posY][posX - 1] != wall {
				cave[posY][posX] = isDiscovered
				posX--
				cave[posY][posX] = player
			}
		case "s":
			if cave[posY + 1][posX] != wall {
				cave[posY][posX] = isDiscovered
				posY++
				cave[posY][posX] = player
			}
		case "d":
			if cave[posY][posX + 1] != wall {
				cave[posY][posX] = isDiscovered
				posX++
				cave[posY][posX] = player
			}
	}
}

func main() {
	generateCave()

	for {
		displayCave()
		move()
	}
}