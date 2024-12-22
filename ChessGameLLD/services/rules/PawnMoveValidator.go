package rules

import (
	"ChessGameLLD/models"
	"errors"
)

type PawnMoveValidator struct {
}

func NewPawnMoveValidator() *PawnMoveValidator {
	return &PawnMoveValidator{}
}

func (pv *PawnMoveValidator) IsValidMove(move *models.Move, piece *models.Piece) error {
	if piece.InitialMove() { //1 or 2 move as initial move
		if move.StartX() == move.EndX() && (move.EndY()-move.StartY() == 1 || move.EndY()-move.StartY() == 2) {
			piece.SetInitialMove(false)
			return nil
		}
	} else {
		if move.StartX() == move.EndX() && (move.EndY()-move.StartY() == 1) {
			return nil
		}
	}
	return errors.New("invalid move for pawn")
}
