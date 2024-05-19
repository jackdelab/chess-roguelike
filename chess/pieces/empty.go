package pieces

import "github.com/jackdelab/chess-roguelike/m/v2/chess"

type EmptyPiece struct{}

var _ chess.Piece = EmptyPiece{}

func (self EmptyPiece) Name() string {
	return ""
}

func (self EmptyPiece) PossibleMoves(int, chess.Coordinate) []chess.Coordinate {
	return []chess.Coordinate{}
}

func (self EmptyPiece) LegalMoves(chess.Board, chess.Coordinate) []chess.Coordinate {
	return []chess.Coordinate{}
}

func (self EmptyPiece) String() string {
	return ""
}

var Empty = EmptyPiece{}
