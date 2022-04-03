package move

type Move struct {
	From, To Position
}

type Position struct {
	File, Rank int
}
