package piece

import (
	"testing"
	"Chess/movement"
	)

func TestKingPosition(t *testing.T){
	king := NewKing(movement.Position{1,1})
	if king.Moves[0].X != -1 || king.Moves[0].Y != -1 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[0])
	}
	if king.Moves[1].X != 0 || king.Moves[1].Y != -1 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[1])
	}
	if king.Moves[2].X != 1 || king.Moves[2].Y != -1 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[2])
	}
	if king.Moves[3].X != 1 || king.Moves[3].Y != 0 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[3])
	}
	if king.Moves[4].X != 1 || king.Moves[4].Y != 1 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[4])
	}
	if king.Moves[5].X != 0 || king.Moves[5].Y != 1 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[5])
	}
	if king.Moves[6].X != -1 || king.Moves[6].Y != 1 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[6])
	}
	if king.Moves[7].X != -1 || king.Moves[7].Y != 0 {
		t.Errorf("Move has wrong position.  Expected: %d",king.Moves[7])
	}
}