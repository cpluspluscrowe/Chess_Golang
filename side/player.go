package side

import (
	"Chess/movement"
	"Chess/color"
)

type Player struct {
	OccupiedPositions     map[movement.Position]bool
	ValidPotentialMoves   map[movement.Position]map[movement.Position]bool // stores potential move, hashset of current positions
	InvalidPotentialMoves map[movement.Position]map[movement.Position]bool
	CheckmateKing         movement.Position
	Color                 color.Color
}

func NewPlayer(color color.Color) *Player {
	player := &Player{}
	player.OccupiedPositions = make(map[movement.Position]bool)
	player.ValidPotentialMoves = make(map[movement.Position]map[movement.Position]bool)
	player.InvalidPotentialMoves = make(map[movement.Position]map[movement.Position]bool)
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

func (player *Player) adjustValidMoves(oldPosition movement.Position, newPosition movement.Position) {
	// if valid positions contains the new position, move them to invalid positions
	if _, ok := player.ValidPotentialMoves[newPosition]; ok {
		player.InvalidPotentialMoves[newPosition] = player.ValidPotentialMoves[newPosition]
		delete(player.ValidPotentialMoves, newPosition)
	}
	// if the old position had invalid moves, then move those positions to valid, since the piece is no longer at that position
	if _, ok := player.InvalidPotentialMoves[oldPosition]; ok {
		player.ValidPotentialMoves[oldPosition] = player.InvalidPotentialMoves[oldPosition]
		delete(player.InvalidPotentialMoves, newPosition)
	}
}

func (player *Player) MovePieceToPosition(oldPosition movement.Position, newPosition movement.Position) {
	_, deleteWillWork := player.OccupiedPositions[oldPosition]
	if !deleteWillWork {
		panic("Old position is not in player's occupied positions.")
	}
	delete(player.OccupiedPositions, oldPosition)
	player.OccupiedPositions[newPosition] = true
	player.adjustValidMoves(oldPosition, newPosition)
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

func (player *Player) SetPotentialMoves() {
	for key, _ := range player.OccupiedPositions {
		player.addMovePotentialPositions(key)
	}
}

func addToPotentialPositionMap(potentialMoveMap map[movement.Position]map[movement.Position]bool,currentPosition movement.Position, newPosition movement.Position){
	if _, ok := potentialMoveMap[currentPosition]; ok {
		potentialMoveMap[currentPosition][newPosition] = true
	}else{
		potentialMoveMap[currentPosition] = make(map[movement.Position]bool)
	}
}

func (player *Player) addToPotentialMovesIfMoveIsValid(currentPosition movement.Position, newPosition movement.Position) {
	if (player.IsMoveValid(newPosition)) {
		addToPotentialPositionMap(player.ValidPotentialMoves, currentPosition, newPosition)
	} else {
		addToPotentialPositionMap(player.InvalidPotentialMoves, currentPosition, newPosition)
	}
}

func (player *Player) addMovePotentialPositions(move movement.Position) {
	x := move.X
	y := move.Y
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x-1, y-1))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x, y-1))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x+1, y-1))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x+1, y))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x+1, y+1))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x, y+1))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x-1, y+1))
	player.addToPotentialMovesIfMoveIsValid(move, movement.NewPosition(x-1, y))
}

func (player *Player) setCheckmateKing(position movement.Position) {
	player.CheckmateKing = position
}
