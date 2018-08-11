package gameboard

import (
	"fmt"
			"Chess/side"
)

type Board struct {
	white *side.Player
	black *side.Player
}

func NewBoard(white *side.Player, black *side.Player) *Board {
	return &Board{white:white, black:black}
}

func (p Board) String() string {
	array := [][]string{
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
	}
	for _, position := range p.black.Pieces {
		array[position.Position.Y][position.Position.X] = "b"
	}
	for _, position := range p.white.Pieces {
		array[position.Position.Y][position.Position.X] = "w"
	}
	var playerString string
	for _, row := range array {
		playerString += row[0] + row[1] + row[2] + row[3] + row[4] + row[5] + row[6] + row[7] + "\n"
	}
	return fmt.Sprintf(playerString)
}

