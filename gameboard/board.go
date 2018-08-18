package gameboard

import (
	"fmt"
			"Chess/side"
	"Chess/movement"
	"Chess/piece"
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
func (b *Board) RemovePiece(position movement.Position, playerTakingPiece *side.Player){
	var sideToRemovePieceFrom *side.Player = b.getOpponent(playerTakingPiece)
	verifyOtherPlayerHasPiece, _ := sideToRemovePieceFrom.OccupiedPositions[position]
	if !verifyOtherPlayerHasPiece {
		panic("No piece at this location to remove!")
	}
	delete(sideToRemovePieceFrom.OccupiedPositions, position)
}

func (b *Board) IsPositionOccupied(position movement.Position, playerTakingPiece *side.Player) bool {
	var sideToRemovePieceFrom *side.Player = b.getOpponent(playerTakingPiece)
	verifyOtherPlayerHasPiece, _ := sideToRemovePieceFrom.OccupiedPositions[position]
	return verifyOtherPlayerHasPiece
}

func (b *Board) MovePiece(player *side.Player, king *piece.King, newPosition movement.Position){
	player.MovePieceToPosition(king, newPosition)
	if b.IsPositionOccupied(newPosition, player){
		b.RemovePiece(newPosition, player)
	}
}

func (b *Board) getOpponent(player *side.Player) *side.Player {
	var opponent *side.Player = nil
	if(player.Color.IsBlack){
		opponent = b.white
	}else{
		opponent = b.black
	}
	return opponent
}

func (b *Board) GetBestMove(player *side.Player){
	opponent := b.getOpponent(player)
	move, king := player.CalculateBestMove(opponent.OccupiedPositions)
	if(b.IsPositionOccupied(move, player)){
		b.RemovePiece(move, player)
	}
	b.MovePiece(player, king, move)
}