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

func NewKing(position Position, color Color) King {
	king := King{}
	king.Moves = []Position{Position{-1,0},
	Position{-1,-1},
		Position{0,-1},
		Position{1,-1},
		Position{1,0},
		Position{1,1},
		Position{0,1},
		Position{-1,1}}
	return king
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

func addRowOfKings(isBlack bool, pieces *[]King){
	var row int
	if isBlack {
		row = 0
	}else{
		row = 7
	}
	for i := 0; i < 8; i++ {
		position := Position{x:i,y:row}
		piece := NewKing(position, Color{isBlack: isBlack})
		*pieces = append(*pieces, piece)
	}
}

type Board struct{
	occupiedPositions map[Position]bool
	pieces []King
}

func NewBoard() Board {
	board := Board{}
	board.occupiedPositions = make(map[Position]bool)
	board.pieces = []King{}
	return board
}

func main(){
	board := NewBoard()


	addRowOfKings(false, &board.pieces)
	addRowOfKings(true, &board.pieces)
	for _, piece := range board.pieces {
		fmt.Println(piece)
	}
}
