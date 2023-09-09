package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"

	"golang.org/x/term"
)

type Cell struct {
	symbol string
	status string
}

const (
	ROWS		  int    = 10
	COLS		  int    = 20
	WALL       string = "🧱"
	DARK       string = "⬛"
	DISCOVERED string = "  "
	PLAYER     string = "😨"
	DIED       string = "😵"
	WON        string = "😁"
	TREASURE   string = "💎"
)

var cave = make([][]Cell, ROWS)
var posX, posY int = 1, 1

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
				
				if rand.Intn(8) == 0 {
					cave[y][x].status = "danger"
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

func playerMove() string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	var move []byte = make([]byte, 1)
	_, err = os.Stdin.Read(move)
	if err != nil {
		panic(err)
	}

	return string(move)
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

func playerDied() {
	cave[posY][posX].symbol = DIED
	displayCave()
	fmt.Println("You just found the death cell 💀")
}

func playerWon() {
	cave[posY][posX].symbol = WON
	displayCave()
	fmt.Println("Congratulations! You have won the game 👏")
}

func resetCells() {
	for y:=1; y<ROWS-1; y++ {
		for x:=1; x<COLS-1; x++ {
			if cave[y][x].symbol == PLAYER || cave[y][x].symbol == TREASURE {
				continue
			}
			cave[y][x].symbol = DARK
		}
	}
}

func shuffleDeathCells() {
	for y:=1; y<ROWS-1; y++ {
		for x:=1; x<COLS-1; x++ {
			if cave[y][x].symbol == PLAYER || cave[y][x].symbol == TREASURE {
				continue
			}
			if rand.Intn(8) == 0 {
				cave[y][x].status = "danger"
			}
		}
	}
}

func timer(done chan bool) {
	for {
		second := 6

		for second >= 0 {
			select {
				case <-done:
					return
				default:
					fmt.Printf("Death cells shuffled in %d\r", second)
					time.Sleep(1 * time.Second)
					second--
			}
		}

		if second == -1 {
			second += 6
			resetCells()
			shuffleDeathCells()
			displayCave()
		}
	}
}

func main() {
	done := make(chan bool)
	go timer(done)
	defer close(done)

	generateCave()

	for {
		displayCave()

		move := playerMove()
		if isNotWall(move) {
			newY, newX := getNewPosition(move)

			cave[posY][posX].symbol = DISCOVERED
			posY, posX = newY, newX

			isDeathCell := cave[newY][newX].status == "danger"
			isTreasure := cave[newY][newX].symbol == TREASURE

			if isDeathCell {
				playerDied()
				break
			} else if isTreasure {
				playerWon()
				break
			}

			cave[posY][posX].symbol = PLAYER
		}
	}
}