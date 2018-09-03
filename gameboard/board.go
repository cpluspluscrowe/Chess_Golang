package gameboard

import (
	"fmt"
			"Chess/side"
	"Chess/movement"
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
	for position, _ := range p.black.OccupiedPositions {
		array[position.Y][position.X] = "b"
	}
	for position, _ := range p.white.OccupiedPositions {
		array[position.Y][position.X] = "w"
	}
	var playerString string
	for _, row := range array {
		playerString += row[0] + row[1] + row[2] + row[3] + row[4] + row[5] + row[6] + row[7] + "\n"
	}
	return fmt.Sprintf(playerString)
}

// will throw exception is there is not a piece to remove
func (b *Board) RemovePiece(position movement.Position, playerLosingPiece *side.Player){
	verifyOtherPlayerHasPiece, _ := playerLosingPiece.OccupiedPositions[position]
	if !verifyOtherPlayerHasPiece {
		panic("No piece at this location to remove!")
	}
	delete(playerLosingPiece.OccupiedPositions, position)
}

func (b *Board) isPositionOccupiedByOpponent(position movement.Position, playerWithPiece *side.Player) bool {
	verifyOtherPlayerHasPiece, _ := playerWithPiece.OccupiedPositions[position]
	return verifyOtherPlayerHasPiece
}

func (b *Board) isPositionAlreadyOccupied(position movement.Position, playerWithPiece *side.Player) bool {
	_, ok := playerWithPiece.OccupiedPositions[position]
	return ok
}

func (b *Board) MovePiece(player *side.Player, oldPosition movement.Position, newPosition movement.Position){
	opponent := b.getOpponent(player)
	playerOccupiesOldPosition := b.isPositionAlreadyOccupied(oldPosition, player)
	playerAlreadyOccupiesNewPosition := b.isPositionAlreadyOccupied(newPosition, player)
	opponentOccupiesNewPosition := b.isPositionAlreadyOccupied(newPosition, opponent)
	if !player.IsMoveValid(newPosition) {
		panic("Move is not valid. New movement: " + fmt.Sprintf("%v",newPosition))
	}
	if !playerOccupiesOldPosition{
		panic("The passed oldPosition is not occupied by a piece.  You cannot move a piece that isn't there" + fmt.Sprintf("%v",oldPosition))
	}
	if playerAlreadyOccupiesNewPosition {
		panic("Position is already occupied by player's pieces.  Occupied position: %d" + fmt.Sprintf("%v",newPosition))
	}
	if opponentOccupiesNewPosition {
		b.RemovePiece(newPosition, opponent)
	}
	player.MovePieceToPosition(oldPosition, newPosition)
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
	oldPosition, newPosition := player.CalculateBestMove(opponent.OccupiedPositions)
	if(b.isPositionOccupiedByOpponent(newPosition, player)){
		b.RemovePiece(newPosition, player)
	}
	b.MovePiece(player, oldPosition, newPosition)
}