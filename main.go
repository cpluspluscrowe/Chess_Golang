package main

import (
	"fmt"
	"Chess/side"
	"Chess/color"
	"Chess/gameboard"
	)

func main(){
	var black = side.NewPlayer(color.Color{true})
	var white = side.NewPlayer(color.Color{false})
	board := gameboard.NewBoard(white, black)
	fmt.Println(white.ValidPotentialMoves)
	fmt.Println(white.InvalidPotentialMoves)
	DoSomeMovement(white, black)
	fmt.Println(board)
}

// todo: fix exception when removing break on line 25, something about trying to move from a position not already occupied
func DoSomeMovement(white *side.Player, black *side.Player){
	for currentPosition,potentialMoves := range white.ValidPotentialMoves {
		for _, potentialMove := range potentialMoves {
			white.MovePieceToPosition(currentPosition,potentialMove)
			break
		}
	}
}
