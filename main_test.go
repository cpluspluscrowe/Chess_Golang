package main

import (
	"testing"
	"fmt"
	"Chess/movement"
	"Chess/side"
	"Chess/color"
)

func TestThatPositionIsOccupiedAfterAddingPiece(t *testing.T){
	player := side.NewPlayer(color.Color{false})
	player.AddKing(movement.Position{1,1}, false)
	ok, _ := player.OccupiedPositions[movement.Position{1,1}]
	if ok {
		fmt.Errorf("movement should be occupied")
	}
}

func TestValidMove(t * testing.T){
	player := side.NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{0,0})
	if !valid {
		fmt.Errorf("Move is valid, but function returned not valid")
	}
}
func TestValidMove2(t * testing.T){
	player := side.NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{7,7})
	if !valid {
		fmt.Errorf("Move is valid, but function returned not valid")
	}
}
func TestValidMove3(t * testing.T){
	player := side.NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{-1,0})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove4(t * testing.T){
	player := side.NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{0,-1})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove5(t * testing.T){
	player := side.NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{8,0})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove6(t * testing.T){
	player := side.NewPlayer(color.Color{false})
	valid := player.IsMoveValid(movement.Position{0,8})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}