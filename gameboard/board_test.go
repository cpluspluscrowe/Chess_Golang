package gameboard

import (
	"testing"
	"Chess/side"
	"Chess/color"
	"Chess/movement"
)

func TestThatBoardIdentifiesThatThePositionIsOccupied(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	shouldBeOccupied := b.isPositionAlreadyOccupied(blackKing, black)
	if !shouldBeOccupied {
		t.Errorf("Position should be occupied.  Occupied Positions for black: ",black.OccupiedPositions)
	}
}

func TestThatTheBoardIdentifiesThatThePositionIsNotOccupied(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	b := NewBoard(white, black)
	shouldBeOccupied := b.isPositionAlreadyOccupied(movement.NewPosition(1,1), black)
	if shouldBeOccupied {
		t.Errorf("Position should not be occupied.  Occupied Positions for black: ",black.OccupiedPositions)
	}
}

func TestThatTheBoardIdentifiesThatThePositionIsNotOccupiedWithOtherPiecesOnTheBoard(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	shouldBeOccupied := b.isPositionAlreadyOccupied(movement.NewPosition(1,1), black)
	if shouldBeOccupied {
		t.Errorf("Position should not be occupied.  Occupied Positions for black: ",black.OccupiedPositions)
	}
}

func TestVerifyThatOldPositionIsOccupied(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	b.isPositionAlreadyOccupied(blackKing, black)
}