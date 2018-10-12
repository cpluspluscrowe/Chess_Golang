package main

import (
	"fmt"
	"Chess/side"
	"Chess/color"
	"Chess/gameboard"
)

// todo: Whites are taking their own pieces.  If the piece is currently occupied by the same side, it should not be a potential move
// todo: Need to update potential moves of moved piece and delete old potential moves for that piece
func main() {
	var black = side.NewPlayer(color.NewColor("black"))
	var white = side.NewPlayer(color.NewColor("white"))
	board := gameboard.NewBoard(white, black)
	DoSomeMovement(white)
	DoSomeMovement(black)
	fmt.Println(board)
}

func DoSomeMovement(white *side.Player) {
	for i := 0; i < 100; i++ {
		for currentPosition, potentialMoves := range white.ValidPotentialMoves {
			for potentialMove,_ := range potentialMoves {
				white.MovePieceToPosition(currentPosition, potentialMove)
				break
			}
		}
	}
}
