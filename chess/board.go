package chess

type DefaultBoard struct {
	height, width int
	squares       []StateCoordinate
}

var _ Board = DefaultBoard{}

func (self DefaultBoard) Dimensions() (int, int) {
	return self.height, self.width
}

func (self DefaultBoard) Piece(Coordinate) Piece
