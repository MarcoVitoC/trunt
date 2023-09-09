package game

import "math/rand"

type Cell struct {
	symbol string
	status string
}

const (
	ROWS       int    = 10
	COLS       int    = 20
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

func GenerateCave() {
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