package rules

import (
	"ChessGameLLD/models"
	"errors"
)

type QueenMoveValidator struct {
	pawnValidator   IRules
	knightValidator IRules
	bishopValidator IRules
	rookValidator   IRules
	kingValidator   IRules
}

func NewQueenMoveValidator() IRules {
	return &QueenMoveValidator{pawnValidator: NewPawnMoveValidator(), knightValidator: NewKnightMoveValidator(),
		bishopValidator: NewBishopMoveValidator(), rookValidator: NewRookMoveValidator(), kingValidator: NewKingMoveValidator()}
}

func (qv *QueenMoveValidator) IsValidMove(move *models.Move, piece *models.Piece) error {
	//Check if move is valid for any of Pawn, Knight, Bishop, Rook, King rules validator
	if qv.pawnValidator.IsValidMove(move, piece) == nil || qv.knightValidator.IsValidMove(move, piece) == nil ||
		qv.bishopValidator.IsValidMove(move, piece) == nil || qv.rookValidator.IsValidMove(move, piece) == nil ||
		qv.kingValidator.IsValidMove(move, piece) == nil {
		return nil
	}

	return errors.New("invalid move for Queeen")
}
