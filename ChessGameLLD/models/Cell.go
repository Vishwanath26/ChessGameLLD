package models

type Cell struct {
	x     int
	y     int
	color string
	piece *Piece
}

func NewCell(x int, y int, color string) *Cell {
	return &Cell{x: x, y: y, color: color}
}

func (c *Cell) X() int {
	return c.x
}

func (c *Cell) Y() int {
	return c.y
}

func (c *Cell) Color() string {
	return c.color
}

func (c *Cell) Piece() *Piece {
	return c.piece
}

func (c *Cell) SetPiece(piece *Piece) {
	c.piece = piece
}
