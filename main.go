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
	blackKing := movement.NewPosition(0,0)
	black.AddKing(blackKing,false)
	board.MovePiece(black, movement.Position{0,0},movement.Position{1,1})
	fmt.Println(board)
}
