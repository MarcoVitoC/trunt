package main

import "fmt"

var cave = make([][]string, 10, 15)

func generateCave() {
	for i:=0; i<10; i++ {
		cave[i] = make([]string, 15)
	}

	for i:=0; i<10; i++ {
		for j:=0; j<15; j++ {
			if i == 0 || i == 10 - 1 || j == 0 || j == 15 - 1 {
				cave[i][j] = "🧱"
			} else {
				cave[i][j] = "⬛"
			}
		}
	}

	cave[1][1] = "🙂"
	cave[8][13] = "💎"

	for i:=0; i<10; i++ {
		for j:=0; j<15; j++ {
			fmt.Print(cave[i][j])
		}
		fmt.Println()
	}
}

func main() {
	generateCave()
}