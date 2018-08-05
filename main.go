package main

import (
	"fmt"
	"math/rand"
)

type Position struct {
	x int
	y int
}

func (player *Player) AddKing(position Position, setAsCheckmateKing bool) {
	player.occupiedPositions[position] = true
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
		player.pieces = append(player.pieces, king)

	if setAsCheckmateKing {
		player.setCheckmateKing(&king)
	}
}

type King struct {
	Position
	Moves []Position
}

func (p King) String() string {
	var color string
	return fmt.Sprintf("{Color:%s, Position:[%d, %d]}", color, p.Position.y, p.Position.x)
}

func isMoveValid(position Position, player *Player) bool {
	if position.x < 0 || position.x > 7 {
		return false
	}
	if position.y < 0 || position.y > 7 {
		return false
	}
	ok, _ := player.occupiedPositions[position]
	if !ok {
		return true
	}else{
		return false
	}
}

func (player *Player) addRowOfKings(row int){
	for i := 0; i < 8; i++ {
		position := Position{x:i,y:row}
		if i == 0 {
			player.AddKing(position, true)
		}else{
			player.AddKing(position, false)
		}
	}
}

type Color struct {
	isBlack bool
}

type Board struct {
	white *Player
	black *Player
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
	for _, position := range p.black.pieces {
		array[position.y][position.x] = "b"
	}
	for _, position := range p.white.pieces {
		array[position.y][position.x] = "w"
	}
	var playerString string
	for _, row := range array {
		playerString += row[0] + row[1] + row[2] + row[3] + row[4] + row[5] + row[6] + row[7] + "\n"
	}
	return fmt.Sprintf(playerString)
}

type PossibleMove struct {
	wouldTakePiece bool
	Position
}

type Player struct{
	occupiedPositions map[Position]bool
	pieces []King
	checkmateKing *King
	Color
	AvailableMoves []PossibleMove
}

func (player *Player) setCheckmateKing(king *King){
	player.checkmateKing = king
}

func NewPlayer(color Color) Player {
	player := Player{}
	player.occupiedPositions = make(map[Position]bool)
	player.pieces = []King{}
	player.checkmateKing = nil
	player.Color = color
	return player
}

func getPieceToMove(player *Player) *King {
	var pieces *[]King
	pieces = &player.pieces
	piece := &(*pieces)[rand.Intn(len(*pieces))]
	return piece
}

func getPieceMove(piece * King, player *Player) (Position, error) {
	for _, move := range piece.Moves {
		moveToPosition := Position{piece.Position.x + move.x, piece.Position.y + move.y}
		if isMoveValid(moveToPosition, player){
			return moveToPosition, nil
		}
	}
	return Position{},fmt.Errorf("No moves are valid for this piece")
}

func movePiece(player *Player){
	blackPiece := getPieceToMove(player)
	wherePieceWillMove, isMoveValid := getPieceMove(blackPiece, player)
	if isMoveValid == nil {
		delete(player.occupiedPositions,blackPiece.Position)
		player.occupiedPositions[wherePieceWillMove] = true
		blackPiece.Position = wherePieceWillMove
	}
}

func main(){
	black := NewPlayer(Color{true})
	white := NewPlayer(Color{false})
	black.addRowOfKings(0)
	white.addRowOfKings(7)
	board := Board{&white, &black}
	fmt.Println(board)
}
