package main

import (
	"testing"
	"fmt"
	"Chess/movement"
	"Chess/side"
	"Chess/color"
)

func TestSetup(t *testing.T){
	white := side.NewPlayer(color.Color{false})
	black := side.NewPlayer(color.Color{true})
	white.AddRowOfKings(0)
	black.AddRowOfKings(7)
	if len(white.Pieces) != 7 {
		fmt.Errorf("lenth should be 7, was equal to: %d", len(white.Pieces))
	}
	if len(black.Pieces) != 7 {
		fmt.Errorf("lenth should be 7, was equal to: %d", len(black.Pieces))
	}
}

func TestThatPositionIsOccupiedAfterAddingPiece(t *testing.T){
	player := side.NewPlayer(color.Color{false})
	player.AddKing(movement.Position{1,1}, false)
	ok, _ := player.OccupiedPositions[movement.Position{1,1}]
	if ok {
		fmt.Errorf("movement should be occupied")
	}
}

func TestThatPositionIsOccupiedAfterAddingPiece2(t *testing.T){
	player := side.NewPlayer(color.Color{false})
	player.AddKing(movement.Position{1,1}, false)
	piece := player.Pieces[0]
	ok, _ := player.OccupiedPositions[piece.Position]
	if ok {
		fmt.Errorf("movement should be occupied")
	}
}

func TestSettingCheckmateKing(t *testing.T){
	player := side.NewPlayer(color.Color{false})
	player.AddKing(movement.Position{1,1}, false)
	king := player.Pieces[0]
	player.SetCheckmateKing(king)
	if player.CheckmateKing == nil {
		fmt.Errorf("Should have set the black check mate king")
	}
}

func TestSettingCheckmateKing2(t *testing.T){
	player := side.NewPlayer(color.Color{false})
	player.AddKing(movement.Position{1,1}, true)
	if player.CheckmateKing == nil {
		fmt.Errorf("Should have set the black check mate king")
	}
	player.AddKing(movement.Position{1,1}, true)
	if player.CheckmateKing == nil {
		fmt.Errorf("Should have set the black check mate king")
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