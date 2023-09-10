package game

func StartGame() {
	done := make(chan bool)
	go Timer(done)
	
	GenerateCave()

	for {
		DisplayCave()

		move := PlayerMove()
		if IsNotWall(move) {
			newY, newX := GetNewPosition(move)

			cave[posY][posX].symbol = DISCOVERED
			posY, posX = newY, newX

			isDeathCell := cave[newY][newX].status == "danger"
			isTreasure := cave[newY][newX].symbol == TREASURE

			if isDeathCell {
				PlayerDied()
				break
			} else if isTreasure {
				PlayerWon()
				break
			}

			cave[posY][posX].symbol = PLAYER
		}
	}

	close(done)
	PlayAgain()
}