package movement

type Position struct {
	X int
	Y int
}

type PossibleMove struct {
	WouldTakePiece bool
	Position Position
}
