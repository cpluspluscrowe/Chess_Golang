package main

import (
	"fmt"
	)

type Position struct {
	x int
	y int
}

type Color struct {
	isBlack bool
}

func AddKing(position Position, color Color, board *Board) {
	king := King{}
	king.Moves = []Position{Position{-1,0},
	Position{-1,-1},
		Position{0,-1},
		Position{1,-1},
		Position{1,0},
		Position{1,1},
		Position{0,1},
		Position{-1,1}}
	if color.isBlack {
		board.blackPieces = append(board.blackPieces, king)
	}else{
		board.whitePieces = append(board.whitePieces, king)
	}
}

type King struct {
	Position
	Color
	Moves []Position
}

func (p King) String() string {
	var color string
	if p.isBlack {
		color = "Black"
	}else{
		color = "White"
	}
	return fmt.Sprintf("{Color:%s, Position:[%d, %d]}", color, p.Position.y, p.Position.x)
}

func isMoveValid(position Position) bool {
	if position.x < 0 || position.x > 7 {
		return false
	}
	if position.y < 0 || position.y > 7 {
		return false
	}
	return true
}

func addRowOfKings(isBlack bool, board *Board){
	var row int
	if isBlack {
		row = 0
	}else{
		row = 7
	}
	for i := 0; i < 8; i++ {
		position := Position{x:i,y:row}
		AddKing(position, Color{isBlack: isBlack}, board)
	}
}

type Board struct{
	occupiedPositions map[Position]bool
	whitePieces []King
	blackPieces []King
}

func NewBoard() Board {
	board := Board{}
	board.occupiedPositions = make(map[Position]bool)
	board.blackPieces = []King{}
	board.whitePieces = []King{}
	return board
}

func main(){
	board := NewBoard()
	addRowOfKings(false, &board)
	addRowOfKings(true, &board)
	fmt.Println(board)
}
