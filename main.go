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
	for currentPosition,potentialMoves := range white.ValidPotentialMoves {
		for _, potentialMove := range potentialMoves {
			white.MovePieceToPosition(currentPosition,potentialMove)
			break
		}
	}
	fmt.Println(board)
}
