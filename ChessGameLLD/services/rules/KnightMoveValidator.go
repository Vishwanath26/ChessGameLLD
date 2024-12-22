package rules

import (
	"ChessGameLLD/models"
	"errors"
	"math"
)

type KnightMoveValidator struct{}

func NewKnightMoveValidator() *KnightMoveValidator {
	return &KnightMoveValidator{}
}

func (kv *KnightMoveValidator) IsValidMove(move *models.Move, piece *models.Piece) error {
	//2 moves in X and 1 moves in Y
	if math.Abs(float64(move.StartX()-move.EndX())) == 2 && math.Abs(float64(move.EndY()-move.EndY())) == 1 {
		return nil
	}
	//Or 2 moves in Y and 1 moves in X
	if math.Abs(float64(move.StartX()-move.EndX())) == 1 && math.Abs(float64(move.EndY()-move.EndY())) == 2 {
		return nil
	}

	return errors.New("invalid move for knight")
}
