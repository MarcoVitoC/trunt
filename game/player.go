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