package game

import (
	"os"

	"golang.org/x/term"
)

func PlayerMove() string {
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

func IsNotWall(move string) bool {
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

func GetNewPosition(move string) (int, int) {
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