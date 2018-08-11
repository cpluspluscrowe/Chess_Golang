package piece

import (
	"Chess/movement"
	"fmt"
	)

type King struct {
	Position movement.Position
	Moves    []movement.Position
}

func (p King) String() string {
	var color string
	return fmt.Sprintf("{color:%s, movement:[%d, %d]}", color, p.Position.Y, p.Position.X)
}

func (king *King) AddPosition(x int, y int) {
	king.Moves = append(king.Moves, movement.Position{x,y})
}

func NewKing(position movement.Position) *King {
	king := &King{}
	king.Position = position
	king.Moves = []movement.Position{}
	king.AddPosition(-1,-1)
	king.AddPosition(0,-1)
	king.AddPosition(1,-1)
	king.AddPosition(1,0)
	king.AddPosition(1,1)
	king.AddPosition(0,1)
	king.AddPosition(-1,1)
	king.AddPosition(-1,0)
	return king
}