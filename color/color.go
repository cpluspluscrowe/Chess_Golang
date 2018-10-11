package color

type Color struct {
	IsBlack bool
}

func NewColor(color string) Color {
	if color == "black" {
		return Color{true}
	}
	if color == "white" {
		return Color{false}
	}
	panic("passed color string should been black or white")
}
