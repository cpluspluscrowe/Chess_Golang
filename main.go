package main

import (
	"fmt"
	"math/rand"
)

type Position struct {
	x int
	y int
}

type Color struct {
	isBlack bool
}

func AddKing(position Position, color Color, board *Board) {
	board.occupiedPositions[position] = true
	king := King{}
	king.Position = position
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

func (p Board) String() string {
	array := [][]string{
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
		{"-","-","-","-","-","-","-","-",},
	}
	for _, position := range p.whitePieces {
		array[position.y][position.x] = "w"
	}
	for _, position := range p.blackPieces {
		array[position.y][position.x] = "b"
	}
	var boardString string
	for _, row := range array {
		boardString += row[0] + row[1] + row[2] + row[3] + row[4] + row[5] + row[6] + row[7] + "\n"
	}
	return fmt.Sprintf(boardString)
}

func getPieceToMove(isBlack bool, board *Board) King {
	var pieces []King
	if isBlack{
		pieces = board.blackPieces
	}else{
		pieces = board.whitePieces
	}
	piece := pieces[rand.Intn(len(pieces))]
	return piece
}

func getPieceMove(piece * King) (Position, error) {
	for _, move := range piece.Moves {
		moveToPosition := Position{piece.Position.x + move.x, piece.Position.y + move.y}
		if isMoveValid(moveToPosition){
			return moveToPosition, nil
		}
	}
	return Position{},nil
}

func main(){
	board := NewBoard()
	addRowOfKings(false, &board)
	addRowOfKings(true, &board)
	/*for i := 0; i < 10 ; i++ {
		blackPiece := getPieceToMove(true, &board)
		wherePieceWillMove, err := getPieceMove(&blackPiece)
		if err != nil {
			delete(board.occupiedPositions,blackPiece.Position)
			board.occupiedPositions[wherePieceWillMove] = true
		}
	}*/
	fmt.Println(board)
}
