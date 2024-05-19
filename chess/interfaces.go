package chess

type Game interface {
	Teams() []string
	NextTeam(string) string
	Player(string) string
	Board() Board
}

type Board interface {
	String() string
	Piece(Coordinate) Piece
	Dimensions() (int, int)
	DropOOB([]Coordinate) []Coordinate
}

type Coordinate interface {
	String() string
	Up() Coordinate
	UpRight() Coordinate
	Right() Coordinate
	DownRight() Coordinate
	Down() Coordinate
	DownLeft() Coordinate
	Left() Coordinate
	UpLeft() Coordinate
}

type Piece interface {
	Name() string
	PossibleMoves(int, Coordinate) []Coordinate
	LegalMoves(Board, Coordinate) []Coordinate
	String() string
}
