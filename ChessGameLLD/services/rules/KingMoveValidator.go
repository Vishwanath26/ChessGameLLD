package rules

import (
	"ChessGameLLD/models"
	"errors"
	"math"
)

type KingMoveValidator struct {
}

func NewKingMoveValidator() IRules {
	return &KingMoveValidator{}
}

func (kv *KingMoveValidator) IsValidMove(move *models.Move, piece *models.Piece) error {
	startX := move.StartX()
	startY := move.StartY()
	endX := move.EndX()
	endY := move.EndY()

	//Allowed 1 move up or down
	if (endX == startX+1 && endY == startY) || (endX == startX-1 && endY == startY) {
		return nil
	}

	//Allowed 1 move left or right
	if (endY == startY+1 && endX == startX) || (endY == startY-1 && endX == startX) {
		return nil
	}

	//Allowed 1 move diagonally
	if math.Abs(float64(move.StartX()-move.EndX())) == 1 &&
		math.Abs(float64(move.StartX()-move.EndX())) == math.Abs(float64(move.EndX()-move.EndY())) {
		return nil
	}

	return errors.New("invalid move for King")
}
