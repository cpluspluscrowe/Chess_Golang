package main

import (
	"fmt"
	"Chess/side"
	"Chess/color"
	"Chess/gameboard"
)

// todo: One a piece moves, update a piece's valid moves, and eliminate its old potential moves from Valid/Invalid moves for that side.Player
func main() {
	var black = side.NewPlayer(color.NewColor("black"))
	var white = side.NewPlayer(color.NewColor("white"))
	board := gameboard.NewBoard(white, black)
	DoSomeMovement(white, black)
	fmt.Println(board)
}

func DoSomeMovement(white *side.Player, black *side.Player) {
	for currentPosition, potentialMoves := range white.ValidPotentialMoves {
		for potentialMove,_ := range potentialMoves {
			white.MovePieceToPosition(currentPosition, potentialMove)
			break
		}
	}
}
