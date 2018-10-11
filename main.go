package main

import (
	"fmt"
	"Chess/side"
	"Chess/color"
	"Chess/gameboard"
	"Chess/movement"
)

func main(){
	var black = side.NewPlayer(color.Color{true})
	var white = side.NewPlayer(color.Color{false})
	board := gameboard.NewBoard(white, black)
	board.MovePiece(white, movement.Position{1,2},movement.Position{2,2})
	fmt.Println(board)
}
