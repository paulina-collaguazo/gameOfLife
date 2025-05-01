package neighbors

func GetCelNeighbors(indexRowBoard, indexColumnBoard int, boardGame, valuRowBoard [][]int) {
	cellNeighborsPosibilites := [][]int{
		{0, 1}, {1, 0}, {1, 1},
		{-1, 0}, {0, -1}, {-1, -1},
		{-1, 1}, {1, -1},
	}
	neighborsValue := make([]int, 0)
	for _, valueNeighbor := range cellNeighborsPosibilites {

		neighborPositionX := indexRowBoard + valueNeighbor[0]
		neighborPositionY := indexColumnBoard + valueNeighbor[1]
		isValidNeighbor := (neighborPositionX >= 0 && neighborPositionX <= len(boardGame) &&
			neighborPositionY >= 0 && neighborPositionY <= len(valuRowBoard))
		if isValidNeighbor {

			neighborsValue = append(neighborsValue, boardGame[neighborPositionX][neighborPositionY])
		}

	}

}
