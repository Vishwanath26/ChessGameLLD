package rules

import (
	"ChessGameLLD/models"
	"errors"
	"math"
)

type RookMoveValidator struct {
}

func NewRookMoveValidator() *RookMoveValidator {
	return &RookMoveValidator{}
}

func (rv *RookMoveValidator) IsValidMove(move *models.Move, piece *models.Piece) error {
	//Any moves in X and 0 moves in Y
	if math.Abs(float64(move.StartX()-move.EndX())) >= 1 && math.Abs(float64(move.EndY()-move.EndY())) == 0 {
		return nil
	}
	//Any moves in Y and 0 moves in X
	if math.Abs(float64(move.StartX()-move.EndX())) == 0 && math.Abs(float64(move.EndY()-move.EndY())) >= 1 {
		return nil
	}

	return errors.New("invalid move for Rook")
}
