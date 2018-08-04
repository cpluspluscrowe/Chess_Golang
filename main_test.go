package main

import (
	"testing"
	"fmt"
)

func TestSetup(t *testing.T){
	pieces := []King{}
	addRowOfKings(false, &pieces)
	addRowOfKings(true, &pieces)
	if len(pieces) != 14 {
		fmt.Errorf("lenth should be 14, was equal to: %d", len(pieces))
	}
}

func TestValidMove(t * testing.T){
	valid := isMoveValid(Position{0,0})
	if !valid {
		fmt.Errorf("Move is valid, but function returned not valid")
	}
}
func TestValidMove2(t * testing.T){
	valid := isMoveValid(Position{7,7})
	if !valid {
		fmt.Errorf("Move is valid, but function returned not valid")
	}
}
func TestValidMove3(t * testing.T){
	valid := isMoveValid(Position{-1,0})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove4(t * testing.T){
	valid := isMoveValid(Position{0,-1})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove5(t * testing.T){
	valid := isMoveValid(Position{8,0})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}
func TestValidMove6(t * testing.T){
	valid := isMoveValid(Position{0,8})
	if valid {
		fmt.Errorf("Move is not valid, but function returned valid")
	}
}