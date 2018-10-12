package side

import (
	"Chess/movement"
	"Chess/color"
)

type Player struct {
	OccupiedPositions     map[movement.Position]bool
	ValidPotentialMoves   map[movement.Position]map[movement.Position]bool // stores potential move, hashset of current positions
	InvalidPotentialMoves map[movement.Position]map[movement.Position]bool
	potentialPositionToCurrentPosition map[movement.Position]map[movement.Position]bool
	CheckmateKing         movement.Position
	Color                 color.Color
}

func NewPlayer(color color.Color) *Player {
	player := &Player{}
	player.OccupiedPositions = make(map[movement.Position]bool)
	player.ValidPotentialMoves = make(map[movement.Position]map[movement.Position]bool)
	player.InvalidPotentialMoves = make(map[movement.Position]map[movement.Position]bool)
	player.potentialPositionToCurrentPosition = make(map[movement.Position]map[movement.Position]bool)
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
	currentPositionsWithPotentialToMoveHere := player.potentialPositionToCurrentPosition[oldPosition] // check for potential moves targeted at the piece's old position
	for currentPositionWithPotentialToMoveHere := range currentPositionsWithPotentialToMoveHere {
		if _, ok := player.InvalidPotentialMoves[currentPositionWithPotentialToMoveHere][newPosition]; !ok {
			panic("The move should exist in player.InvalidPotentialMoves, since the old position was filled")
			delete(player.InvalidPotentialMoves[currentPositionWithPotentialToMoveHere],newPosition)
		}
		player.ValidPotentialMoves[currentPositionWithPotentialToMoveHere][oldPosition] = true
	}
	currentPositionOfPiecesTargetingWhereTheNewPieceWasMoved := player.potentialPositionToCurrentPosition[newPosition]
	for currentPiecePosition := range currentPositionOfPiecesTargetingWhereTheNewPieceWasMoved {
		if _, ok := player.ValidPotentialMoves[currentPiecePosition][newPosition]; !ok {
			panic("The move should exist in player.ValidPotentialMoves, since the new position was not taken until just now")
			delete(player.InvalidPotentialMoves[currentPiecePosition],newPosition)
		}
		player.InvalidPotentialMoves[currentPiecePosition][newPosition] = false
	}
	// now that the piece has moved, there are no potential moves that begin at the old position.
	delete(player.ValidPotentialMoves,oldPosition)
	// now recalcualte potential moves from the piece's new position
	player.addMovePotentialPositions(newPosition)
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

func (player *Player) addToPotentialPositionMap(potentialMoveMap map[movement.Position]map[movement.Position]bool,currentPosition movement.Position, newPosition movement.Position, isValid bool){
	if _, ok := potentialMoveMap[currentPosition]; ok {
		potentialMoveMap[currentPosition][newPosition] = true
		player.potentialPositionToCurrentPosition[newPosition][currentPosition] = isValid
	}else{
		potentialMoveMap[currentPosition] = make(map[movement.Position]bool)
		player.potentialPositionToCurrentPosition[newPosition] = make(map[movement.Position]bool)
	}
}

func (player *Player) addToPotentialMovesIfMoveIsValid(currentPosition movement.Position, newPosition movement.Position) {
	if player.IsMoveValid(newPosition) {
		player.addToPotentialPositionMap(player.ValidPotentialMoves, currentPosition, newPosition, true)
	} else {
		player.addToPotentialPositionMap(player.InvalidPotentialMoves, currentPosition, newPosition, false)
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
