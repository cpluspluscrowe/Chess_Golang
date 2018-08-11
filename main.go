package main

import (
	"fmt"
	"math/rand"
	"Chess/piece"
	"Chess/side"
	"Chess/color"
	"Chess/gameboard"
)



func getPieceToMove(player *side.Player) *piece.King {
	var pieces *[]*piece.King
	pieces = &player.Pieces
	piece := (*pieces)[rand.Intn(len(*pieces))]
	return piece
}

func main(){
	var black = side.NewPlayer(color.Color{true})
	var white = side.NewPlayer(color.Color{false})
	board := gameboard.NewBoard(white, black)
	black.AddRowOfKings(0)
	white.AddRowOfKings(7)
	fmt.Println(board)

	black.MovePiece(black.Pieces[0],1,1)

	fmt.Println(board)
}
