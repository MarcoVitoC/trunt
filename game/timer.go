package game

import (
	"fmt"
	"time"
)

func Timer(done chan bool) {
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
			ResetCells()
			ShuffleDeathCells()
			DisplayCave()
		}
	}
}