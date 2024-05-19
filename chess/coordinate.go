package chess

import (
	"strconv"
	"strings"
)

func invertColor(in Color) Color {
	switch in {
	case Black:
		return White
	case White:
		return Black
	default:
		return OOB
	}
}

type Square struct {
	color       Color
	column, row int
}

var _ Coordinate = Square{}

func (self Square) String() string {
	columnOffset := 97
	rowOffset := 1

	columnStr := string(rune(self.column + columnOffset))
	rowStr := strconv.Itoa(self.row + rowOffset)

	return strings.Join([]string{columnStr, rowStr}, "")
}

func (self Square) Up() Coordinate {
	return Square{
		color:  invertColor(self.color),
		column: self.column,
		row:    self.row + 1,
	}
}

func (self Square) Right() Coordinate {
	return Square{
		color:  invertColor(self.color),
		column: self.column + 1,
		row:    self.row,
	}
}

func (self Square) Down() Coordinate {
	return Square{
		color:  invertColor(self.color),
		column: self.column,
		row:    self.row - 1,
	}
}

func (self Square) Left() Coordinate {
	return Square{
		color:  invertColor(self.color),
		column: self.column - 1,
		row:    self.row,
	}
}

func (self Square) UpRight() Coordinate {
	return self.Up().Right()
}

func (self Square) UpLeft() Coordinate {
	return self.Up().Left()
}

func (self Square) DownRight() Coordinate {
	return self.Down().Right()
}

func (self Square) DownLeft() Coordinate {
	return self.Down().Left()
}
