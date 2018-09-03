package side

import (
	"Chess/movement"
	"Chess/color"
	"fmt"
)

type Player struct{
	OccupiedPositions map[movement.Position]bool
	CheckmateKing movement.Position
	Color color.Color
}

func (player *Player) SetCheckmateKing(position movement.Position){
	player.CheckmateKing = position
}

func NewPlayer(color color.Color) *Player {
	player := &Player{}
	player.OccupiedPositions = make(map[movement.Position]bool)
	player.CheckmateKing = movement.Position{}
	player.Color = color
	return player
}

func (player *Player) AddKing(position movement.Position, setAsCheckmateKing bool) {
	player.OccupiedPositions[position] = true
	if setAsCheckmateKing {
		player.SetCheckmateKing(position)
	}
}
func (player *Player) MovePieceXY(oldPosition movement.Position,xChange int, yChange int){
	newPosition := movement.Position{oldPosition.X + xChange, oldPosition.Y + yChange}
	if !player.IsMoveValid(newPosition) {
		fmt.Errorf("Move is not valid. New movement: %d", newPosition)
	}
}

func (player *Player) MovePieceToPosition(oldPosition movement.Position,newPosition movement.Position){
	_, deleteWillWork := player.OccupiedPositions[oldPosition]
	if !deleteWillWork {
		panic("Old position is not in player's occupied positions.")
	}
	delete(player.OccupiedPositions,oldPosition)
	player.OccupiedPositions[newPosition] = true
	fmt.Println(player.OccupiedPositions)
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
	if ok {
		return false
	}
	return true
}

// TODO: FIX ME!  YOU NEED TO PROVIDE VALID POSITIONS FOR EACH OCCUPIED POSITION
func (player *Player) CalculateBestMove(movesOccupiedByOtherColor map[movement.Position]bool) (movement.Position, movement.Position){
	var bestMove movement.Position
		for move, _ := range player.OccupiedPositions {
			if player.IsMoveValid(move) {
				panic("This needs to be fixed, returning two of the same moves right now")
				return move,move
			}
		}
	fmt.Errorf("Player is unable to move!  Player: %d", player.Color)
	return bestMove, bestMove
}