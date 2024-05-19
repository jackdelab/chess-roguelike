package pieces

import (
	"github.com/jackdelab/chess-roguelike/m/v2/chess"
	"github.com/jackdelab/chess-roguelike/m/v2/utils"
)

type Pawn struct {
	color chess.Color
}

var _ chess.Piece = Pawn{}

func (self Pawn) Name() string {
	return "Pawn"
}

func (self Pawn) String() string {
	switch self.color {
	case chess.Black:
		return "♟︎"
	case chess.White:
		return "♙"
	default:
		return "p"
	}
}

func (self Pawn) PossibleMoves(maxDimension int, start chess.Coordinate) []chess.Coordinate {
	var _ = maxDimension // unused for Pawns
	destinations := []chess.Coordinate{}

	if self.color == chess.White {
		// Up is only forward for White
		destinations = append(destinations, start.UpLeft(), start.Up(), start.UpRight())
	} else {
		destinations = append(destinations, start.DownLeft(), start.Down(), start.DownRight())
	}

	return destinations
}

func (self Pawn) LegalMoves(board chess.Board, start chess.Coordinate) []chess.Coordinate {
	maxDimension := utils.Max(board.Dimensions())
	moves := self.PossibleMoves(maxDimension, start)
	destinations := []chess.Coordinate{}

	if self.color == chess.White {
		for _, move := range moves {
			if move == start.Up() && board.Piece(move) == Empty {
				destinations = append(destinations, move)
			} else if move != start.Up() && board.Piece(move) != Empty {
				destinations = append(destinations, move)
			}
		}
	} else {
		for _, move := range moves {
			if move == start.Down() && board.Piece(move) == Empty {
				destinations = append(destinations, move)
			} else if move != start.Down() && board.Piece(move) != Empty {
				destinations = append(destinations, move)
			}
		}
	}

	return board.DropOOB(destinations)
}
