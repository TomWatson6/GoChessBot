package chess

type Move struct {
	From, To Position
}

type Position struct {
	File, Rank byte
}
