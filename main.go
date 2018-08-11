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
	black.AddRowOfKings(0)
	white.AddRowOfKings(7)
	fmt.Println(board)

	black.MovePieceXY(black.Pieces[0],1,1)

	fmt.Println(board)
}
