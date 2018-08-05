package main

import (
	"testing"
	"fmt"
)

func TestSetup(t *testing.T){
	player := NewPlayer()
	addRowOfKings(false, &player)
	addRowOfKings(true, &player)
	if len(player.pieces) != 7 {
		fmt.Errorf("lenth should be 7, was equal to: %d", len(player.pieces))
	}
	if len(player.pieces) != 7 {
		fmt.Errorf("lenth should be 7, was equal to: %d", len(player.pieces))
	}
}

func TestThatPositionIsOccupiedAfterAddingPiece(t *testing.T){
	player := NewPlayer()
	player.AddKing(Position{1,1}, false)
	ok, _ := player.occupiedPositions[Position{1,1}]
	if ok {
		fmt.Errorf("Position should be occupied")
	}
}

func TestThatPositionIsOccupiedAfterAddingPiece2(t *testing.T){
	player := NewPlayer()
	player.AddKing(Position{1,1}, false)
	piece := player.pieces[0]
	ok, _ := player.occupiedPositions[piece.Position]
	if ok {
		fmt.Errorf("Position should be occupied")
	}
}

func TestSettingCheckmateKing(t *testing.T){
	player := NewPlayer()
	player.AddKing(Position{1,1}, false)
	king := &player.pieces[0]
	player.setCheckmateKing(king)
	if player.checkmateKing == nil {
		fmt.Errorf("Should have set the black check mate king")
	}
}

func TestSettingCheckmateKing2(t *testing.T){
	player := NewPlayer()
	player.AddKing(Position{1,1}, true)
	if player.checkmateKing == nil {
		fmt.Errorf("Should have set the black check mate king")
	}
	player.AddKing(Position{1,1}, true)
	if player.checkmateKing == nil {
		fmt.Errorf("Should have set the black check mate king")
	}
}

func TestValidMove(t * testing.T){
	player := NewPlayer()
	valid := isMoveValid(Position{0,0}, &player)
	if !valid {
		fmt.Errorf("Move is valid, but function returned not valid")
	}
}
func TestValidMove2(t * testing.T){
	player := NewPlayer()
	valid := isMoveValid(Position{7,7}, &player)
	if !valid {
		fmt.Errorf("Move is valid, but function returned not valid")
	}
}
func TestValidMove3(t * testing.T){
	player := NewPlayer()
	valid := isMoveValid(Position{-1,0}, &player)
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove4(t * testing.T){
	player := NewPlayer()
	valid := isMoveValid(Position{0,-1}, &player)
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove5(t * testing.T){
	player := NewPlayer()
	valid := isMoveValid(Position{8,0}, &player)
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove6(t * testing.T){
	player := NewPlayer()
	valid := isMoveValid(Position{0,8}, &player)
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}