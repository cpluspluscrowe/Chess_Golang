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
	potentialWhiteMoves := white.GetPotentialMoves()
	fmt.Println(potentialWhiteMoves)
	for currentPosition,potentialMoves := range potentialWhiteMoves {
		for potentialMove,_ := range potentialMoves {
			white.MovePieceToPosition(currentPosition,potentialMove)
			break
		}
	}
	fmt.Println(board)
}
