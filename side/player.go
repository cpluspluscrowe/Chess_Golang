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

func (player *Player) GetPotentialMoves() map[movement.Position]map[movement.Position]bool {
	potentialPositions := make(map[movement.Position]map[movement.Position]bool)
	for key, _ := range player.OccupiedPositions {
		potentialPositions = player.addMovePotentialPositions(key, &potentialPositions)
		fmt.Println(len(potentialPositions))
	}
	return potentialPositions
}

// store if the move is valid as the final value
// the goal is to update the validity of all positions with each move
// This means that we want to keep a map of position -> list of pieces so I can change move validity
// I'll want two maps.  One with valid moves and one with invalid moves.  I can shuffle between valid and invalid with each move
// makes each move a constant time change to keep a map of possible moves on the board!  Not O(n), since that much interference is impossible, maybe ~4 shuffles per change
func (player *Player) addToPotentialMovesIfMoveIsValid(currentPosition movement.Position,newPosition movement.Position, potentialPositions *map[movement.Position]map[movement.Position]bool){
	isValid := player.IsMoveValid(newPosition)
	// if the inner hashset does not exist, create the map before populating it
	if _, ok := (*potentialPositions)[currentPosition]; !ok {
		(*potentialPositions)[currentPosition] = make(map[movement.Position]bool)
	}
	(*potentialPositions)[currentPosition][newPosition] = isValid
}

func (player *Player) addMovePotentialPositions(move movement.Position, potentialPositions *map[movement.Position]map[movement.Position]bool) map[movement.Position]map[movement.Position]bool {
	x := move.X
	y := move.Y
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x-1,y-1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x,y-1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x+1,y-1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x+1,y),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x+1,y+1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x,y+1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x-1,y+1),potentialPositions)
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x-1,y),potentialPositions)
	return *potentialPositions
}

func (player *Player) setCheckmateKing(position movement.Position){
	player.CheckmateKing = position
}
