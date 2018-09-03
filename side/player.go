package side

import (
	"Chess/movement"
	"Chess/color"
)

type Player struct{
	OccupiedPositions map[movement.Position]bool
	CheckmateKing movement.Position
	Color color.Color
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
		player.setCheckmateKing(position)
	}
}

func (player *Player) MovePieceToPosition(oldPosition movement.Position,newPosition movement.Position){
	_, deleteWillWork := player.OccupiedPositions[oldPosition]
	if !deleteWillWork {
		panic("Old position is not in player's occupied positions.")
	}
	delete(player.OccupiedPositions,oldPosition)
	player.OccupiedPositions[newPosition] = true
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

func (player *Player) GetPotentialMoves() map[movement.Position]bool {
	potentialPositions := make(map[movement.Position]bool)
	for key, _ := range player.OccupiedPositions {
		potentialPositions = player.addMovePotentialPositions(key, &potentialPositions)
	}
	return potentialPositions
}

func (player *Player) addToPotentialMovesIfMoveIsValid(move movement.Position, potentialPositions *map[movement.Position]bool){
	if player.IsMoveValid(move){
		(*potentialPositions)[move] = true
	}
}

func (player *Player) addMovePotentialPositions(move movement.Position, potentialPositions *map[movement.Position]bool) map[movement.Position]bool{
	x := move.X
	y := move.Y
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x-1,y-1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x,y-1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x+1,y-1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x+1,y),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x+1,y+1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x,y+1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x-1,y+1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(movement.NewPosition(x-1,y),potentialPositions)
	return *potentialPositions
}

func (player *Player) setCheckmateKing(position movement.Position){
	player.CheckmateKing = position
}
