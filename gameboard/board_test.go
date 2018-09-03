package gameboard

import (
	"testing"
	"Chess/side"
	"Chess/color"
	"Chess/movement"
	"fmt"
)

func TestThatBoardIdentifiesThatThePositionIsOccupied(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	shouldBeOccupied := b.isPositionAlreadyOccupied(blackKing, black)
	if !shouldBeOccupied {
		t.Errorf("Position should be occupied.  Occupied Positions for black: " + fmt.Sprintf("%v",black.OccupiedPositions))
	}
}

func TestThatTheBoardIdentifiesThatThePositionIsNotOccupied(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	b := NewBoard(white, black)
	shouldBeOccupied := b.isPositionAlreadyOccupied(movement.NewPosition(1,1), black)
	if shouldBeOccupied {
		t.Errorf("Position should not be occupied.  Occupied Positions for black: " + fmt.Sprintf("%v",black.OccupiedPositions))
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
		t.Errorf("Position should not be occupied.  Occupied Positions for black: " + fmt.Sprintf("%v",black.OccupiedPositions))
	}
}

func TestRemovePiece(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	b.RemovePiece(blackKing, black)
	shouldNotBeOccupied := b.isPositionAlreadyOccupied(movement.NewPosition(1,1), black)
	if shouldNotBeOccupied {
		t.Errorf("Position should not be occupied.  Occupied Positions for black: " + fmt.Sprintf("%v",black.OccupiedPositions))
	}
}

func TestMovePiece(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	newPosition := movement.NewPosition(1,1)
	b.MovePiece(black, blackKing, newPosition)
	//b.RemovePiece(blackKing, black)
	shouldBeOccupied := b.isPositionAlreadyOccupied(movement.NewPosition(1,1), black)
	shouldNotBeOccupied := b.isPositionAlreadyOccupied(movement.NewPosition(0,0), black)
	if !shouldBeOccupied || shouldNotBeOccupied {
		t.Errorf("Position should be occupied.  Occupied Positions for black: " + fmt.Sprintf("%v",black.OccupiedPositions))
	}
}

func TestBoardToString(t *testing.T){
	white := side.NewPlayer(color.NewColor(false))
	black := side.NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	b := NewBoard(white, black)
	expected := `b-------
--------
--------
--------
--------
--------
--------
--------
`
	if b.String() != expected {
		t.Errorf("Wrong String representation of the board: " + fmt.Sprintf("%s",b.String()) + "\n" + expected)
	}
}