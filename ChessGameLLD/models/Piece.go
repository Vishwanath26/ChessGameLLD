package models

import "github.com/google/uuid"

const (
	Pawn   = "pawn"
	Rook   = "rook"
	Bishop = "bishop"
	Knight = "knight"
	King   = "king"
	Queen  = "queen"
)

type Piece struct {
	id          string
	name        string
	initialMove bool
}

func NewPiece(name string) *Piece {
	return &Piece{
		id:          uuid.NewString(),
		name:        name,
		initialMove: true,
	}
}

func (p *Piece) InitialMove() bool {
	return p.initialMove
}

func (p *Piece) SetInitialMove(initialMove bool) {
	p.initialMove = initialMove
}

func (p *Piece) Id() string {
	return p.id
}

func (p *Piece) Name() string {
	return p.name
}
