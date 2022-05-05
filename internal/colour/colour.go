package colour

type Colour byte

const (
	White Colour = iota
	Black
)

func (c Colour) String() string {
	switch c {
	case White:
		return "White"
	case Black:
		return "Black"
	default:
		return "Unknown"
	}
}
