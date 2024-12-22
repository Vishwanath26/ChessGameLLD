package rules

import (
	"ChessGameLLD/models"
	"errors"
	"math"
)

type BishopMoveValidator struct {
}

func NewBishopMoveValidator() *BishopMoveValidator {
	return &BishopMoveValidator{}
}

func (bv *BishopMoveValidator) IsValidMove(move *models.Move, piece *models.Piece) error {
	if math.Abs(float64(move.StartX()-move.EndX())) == math.Abs(float64(move.EndY()-move.EndY())) { //1 move in any diagonal
		return nil
	}
	return errors.New("invalid move for bishop")
}
