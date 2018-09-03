package side

import (
	"testing"
	"Chess/color"
	"Chess/movement"
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