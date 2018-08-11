package side

import (
	"Chess/movement"
	"Chess/piece"
	"Chess/color"
	"fmt"
)

type Player struct{
	OccupiedPositions map[movement.Position]bool
	Pieces []*piece.King
	CheckmateKing *piece.King
	Color color.Color
	AvailableMoves []movement.PossibleMove
}

func (player *Player) SetCheckmateKing(king *piece.King){
	player.CheckmateKing = king
}

func NewPlayer(color color.Color) *Player {
	player := &Player{}
	player.OccupiedPositions = make(map[movement.Position]bool)
	player.Pieces = []*piece.King{}
	player.CheckmateKing = nil
	player.Color = color
	return player
}

func (player *Player) AddKing(position movement.Position, setAsCheckmateKing bool) {
	player.OccupiedPositions[position] = true
	king := piece.NewKing(position)
	king.Position = position
	player.Pieces = append(player.Pieces, king)
	if setAsCheckmateKing {
		player.SetCheckmateKing(king)
	}
}
func (player *Player) MovePiece(king *piece.King,xChange int, yChange int){
	king.Position.X = king.Position.X + xChange
	king.Position.Y = king.Position.Y + yChange
	if !player.IsMoveValid(king.Position) {
		fmt.Errorf("Move is not valid. New movement: %d", king.Position)
	}
}

func (player *Player) AddRowOfKings(row int){
	for i := 0; i < 8; i++ {
		position := movement.Position{X: i,Y:row}
		if i == 0 {
			player.AddKing(position, true)
		}else{
			player.AddKing(position, false)
		}
	}
}

func (player *Player) IsMoveValid(position movement.Position) bool {
	if position.X < 0 || position.X > 7 {
		return false
	}
	if position.Y < 0 || position.Y > 7 {
		return false
	}
	ok, _ := player.OccupiedPositions[position]
	if !ok {
		return true
	}else{
		return false
	}
}