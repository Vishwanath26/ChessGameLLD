package rules

import "ChessGameLLD/models"

type IRules interface {
	IsValidMove(move *models.Move, piece *models.Piece) error
}

type Validator struct {
	pawnValidator   IRules
	knightValidator IRules
	bishopValidator IRules
	rookValidator   IRules
	kingValidator   IRules
	queenValidator  IRules
}

func NewRulesValidator() *Validator {
	validator := &Validator{
		pawnValidator:   NewPawnMoveValidator(),
		knightValidator: NewKnightMoveValidator(),
		bishopValidator: NewBishopMoveValidator(),
		rookValidator:   NewRookMoveValidator(),
		kingValidator:   NewKingMoveValidator(),
		queenValidator:  NewQueenMoveValidator(),
	}
	return validator
}

func (rv *Validator) ValidateRules(move *models.Move, piece *models.Piece) error {
	switch piece.Name() {
	case models.Pawn:
		err := rv.PawnValidator().IsValidMove(move, piece)
		if err != nil {
			return err
		}
		break

	case models.Rook:
		err := rv.RookValidator().IsValidMove(move, piece)
		if err != nil {
			return err
		}
		break

	case models.Knight:
		err := rv.KnightValidator().IsValidMove(move, piece)
		if err != nil {
			return err
		}
		break

	case models.Bishop:
		err := rv.BishopValidator().IsValidMove(move, piece)
		if err != nil {
			return err
		}
		break

	case models.Queen:
		err := rv.QueenValidator().IsValidMove(move, piece)
		if err != nil {
			return err
		}
		break

	case models.King:
		err := rv.KingValidator().IsValidMove(move, piece)
		if err != nil {
			return err
		}
	}
	return nil
}

func (rv *Validator) PawnValidator() IRules {
	return rv.pawnValidator
}

func (rv *Validator) KnightValidator() IRules {
	return rv.knightValidator
}

func (rv *Validator) BishopValidator() IRules {
	return rv.bishopValidator
}

func (rv *Validator) RookValidator() IRules {
	return rv.rookValidator
}

func (rv *Validator) KingValidator() IRules {
	return rv.kingValidator
}

func (rv *Validator) QueenValidator() IRules {
	return rv.queenValidator
}
