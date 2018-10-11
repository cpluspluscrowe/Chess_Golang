package color

type Color struct {
	IsBlack bool
}

func NewColor(isBlack bool) Color {
	return Color{isBlack}
}
