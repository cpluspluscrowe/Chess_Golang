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

func (board *Board) AddKing(position Position, color Color, setAsCheckmateKing bool) {
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
	if setAsCheckmateKing {
		board.setCheckmateKing(color, &king)
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

func isMoveValid(position Position, board *Board) bool {
	if position.x < 0 || position.x > 7 {
		return false
	}
	if position.y < 0 || position.y > 7 {
		return false
	}
	ok, _ := board.occupiedPositions[position]
	if !ok {
		return true
	}else{
		return false
	}
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
		if i == 0 {
			board.AddKing(position, Color{isBlack: isBlack}, true)
		}else{
			board.AddKing(position, Color{isBlack: isBlack}, false)
		}
	}
}

type Board struct{
	occupiedPositions map[Position]bool
	whitePieces []King
	blackPieces []King
	whiteCheckmateKing *King
	blackCheckmateKing *King
}

func (board *Board) setCheckmateKing(color Color, king *King){
	if king.Color.isBlack {
		board.blackCheckmateKing = king
	}else{
		board.whiteCheckmateKing = king
	}
}

func NewBoard() Board {
	board := Board{}
	board.occupiedPositions = make(map[Position]bool)
	board.blackPieces = []King{}
	board.whitePieces = []King{}
	board.whiteCheckmateKing = nil
	board.blackCheckmateKing = nil
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

func getPieceToMove(isBlack bool, board *Board) *King {
	var pieces *[]King
	if isBlack{
		pieces = &board.blackPieces
	}else{
		pieces = &board.whitePieces
	}
	piece := &(*pieces)[rand.Intn(len(*pieces))]
	return piece
}

func getPieceMove(piece * King, board *Board) (Position, error) {
	for _, move := range piece.Moves {
		moveToPosition := Position{piece.Position.x + move.x, piece.Position.y + move.y}
		if isMoveValid(moveToPosition, board){
			return moveToPosition, nil
		}
	}
	return Position{},fmt.Errorf("No moves are valid for this piece")
}

func movePiece(color Color, board *Board){
	blackPiece := getPieceToMove(color.isBlack, board)
	wherePieceWillMove, isMoveValid := getPieceMove(blackPiece, board)
	if isMoveValid == nil {
		delete(board.occupiedPositions,blackPiece.Position)
		board.occupiedPositions[wherePieceWillMove] = true
		blackPiece.Position = wherePieceWillMove
	}
}

func main(){
	board := NewBoard()
	addRowOfKings(false, &board)
	addRowOfKings(true, &board)
	for i := 0; i < 1000 ; i++ {
		movePiece(Color{true}, &board)
		movePiece(Color{false}, &board)
	}
	fmt.Println(board)
}
