package gameboard

import (
	"fmt"
			"Chess/side"
	"Chess/movement"
	"Chess/color"
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

// will throw exception is there is not a piece to remove
func (b *Board) RemovePiece(position movement.Position, colorTakingPiece color.Color){
	var sideToRemovePieceFrom *side.Player = nil
	if(b.white.Color == colorTakingPiece){
		sideToRemovePieceFrom = b.black
	}else if(b.black.Color == colorTakingPiece){
		sideToRemovePieceFrom = b.white
	}else{
		panic("Given color did not match either piece")
	}
	verifyOtherPlayerHasPiece, _ := sideToRemovePieceFrom.OccupiedPositions[position]
	if !verifyOtherPlayerHasPiece {
		panic("No piece at this location to remove!")
	}
	delete(sideToRemovePieceFrom.OccupiedPositions, position)
}

func (b *Board) IsPositionOccupied(position movement.Position, colorTakingPiece color.Color) bool {
	var sideToRemovePieceFrom *side.Player = nil
	if(b.white.Color == colorTakingPiece){
		sideToRemovePieceFrom = b.black
	}else if(b.black.Color == colorTakingPiece){
		sideToRemovePieceFrom = b.white
	}else{
		panic("Given color did not match either piece")
	}
	verifyOtherPlayerHasPiece, _ := sideToRemovePieceFrom.OccupiedPositions[position]
	return verifyOtherPlayerHasPiece
}