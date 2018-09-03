package side

import (
	"testing"
	"Chess/color"
	"Chess/movement"
	"fmt"
)

func TestVerifyThatOldPositionIsOccupied(t *testing.T){
	black := NewPlayer(color.NewColor(true))
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing, false)
	potentialPositions := make(map[movement.Position]bool)
	returned := black.addMovePotentialPositions(blackKing, &potentialPositions)
	if len(returned) != 3 {
		t.Errorf("Should return three positions")
	}
}

func TestThatPositionIsOccupiedAfterAddingPiece(t *testing.T){
	player := NewPlayer(color.Color{false})
	player.AddKing(movement.Position{1,1}, false)
	ok, _ := player.OccupiedPositions[movement.Position{1,1}]
	if ok {
		fmt.Errorf("movement should be occupied")
	}
}

func TestValidMove(t *testing.T) {
	player := NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{0,8})
	if valid {
		t.Errorf("Move is not valid, but function returned valid")
	}
	valid = player.IsMoveValid(movement.Position{8,0})
	if valid {
		t.Errorf("Move is not valid, but function returned valid")
	}
	valid = player.IsMoveValid(movement.Position{0,-1})
	if valid {
		t.Errorf("Move is not valid, but function returned valid")
	}
	valid = player.IsMoveValid(movement.Position{-1,0})
	if valid {
		t.Errorf("Move is not valid, but function returned valid")
	}
	valid = player.IsMoveValid(movement.Position{7,7})
	if !valid {
		t.Errorf("Move is valid, but function returned not valid")
	}
	valid = player.IsMoveValid(movement.Position{0,0})
	if !valid {
		t.Errorf("Move is valid, but function returned not valid")
	}
}