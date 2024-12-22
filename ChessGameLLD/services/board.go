package services

import (
	"ChessGameLLD/models"
	"ChessGameLLD/services/rules"
	"errors"
)

const (
	InProgress = "inProgress"
	Finished   = "finished"
)

type BoardService struct {
	PlayerService  IPlayerService
	BoardMapping   map[string]*models.Board
	RulesValidator *rules.Validator
}

func NewBoardService(service IPlayerService) IBoardService {
	return &BoardService{
		PlayerService:  service,
		BoardMapping:   make(map[string]*models.Board),
		RulesValidator: rules.NewRulesValidator(),
	}
}

var COLOURS = []string{"black", "white"}

type IBoardService interface {
	CreateBoard(size int) string
	PlacePieceOnBoard(boardID string, piece *models.Piece, x int, y int) error
	ValidateMove(boardID string, playerID string, pieceName string, move *models.Move) error
	MakeMove(boardID string, playerID string, pieceName string, move *models.Move) error
	CheckGameStatus(boardID string, playerID string) (string, *models.Player)
	isMoveValid(board *models.Board, move *models.Move) error
}

func (bs *BoardService) CreateBoard(size int) string {
	cells := make([][]models.Cell, size)
	count := 0

	for i := 0; i < size; i++ {
		cells[i] = make([]models.Cell, size)
		for j := 0; j < size; j++ {
			cells[i][j] = *models.NewCell(i, j, COLOURS[count%2])
			count++
		}
	}

	newBoard := models.NewBoard(cells)
	bs.BoardMapping[newBoard.Id()] = newBoard
	return newBoard.Id()
}

func (bs *BoardService) PlacePieceOnBoard(boardID string, piece *models.Piece, x int, y int) error {
	board := bs.BoardMapping[boardID]
	cell := board.Cells()[x][y]
	if cell.Piece() != nil {
		return errors.New("cell already has a piece")
	}
	cell.SetPiece(piece)
	return nil
}

func (bs *BoardService) ValidateMove(boardID string, playerID string, pieceName string, move *models.Move) error {
	//Basic validation of params
	if bs.BoardMapping[boardID] == nil {
		return errors.New("invalid board id")
	}
	if bs.PlayerService.PlayersMapping()[playerID] == nil {
		return errors.New("invalid player id")
	}

	board := bs.BoardMapping[boardID]
	startCell := board.Cells()[move.StartX()][move.StartY()]
	endCell := board.Cells()[move.EndX()][move.EndY()]

	currentPlayer := bs.PlayerService.PlayersMapping()[playerID]
	opponent := bs.PlayerService.PlayersMapping()[board.Player1ID()]
	if currentPlayer.Id() == board.Player1ID() {
		opponent = bs.PlayerService.PlayersMapping()[board.Player2ID()]
	}

	var currentPiece *models.Piece
	for _, piece := range currentPlayer.Pieces() {
		if piece.Name() == pieceName {
			currentPiece = &piece
			break
		}
	}

	//validate the move coordinates (within board)
	err := bs.isMoveValid(board, move)
	if err != nil {
		return err
	}

	//Check if start cell has current player's piece and end cell has opponent's piece
	err = checkIfCorrectPiecesPresentInMoveCells(currentPlayer, opponent, startCell, endCell, pieceName)
	if err != nil {
		return err
	}

	//Validation according to Piece movement rules
	err = bs.RulesValidator.ValidateRules(move, currentPiece)
	if err != nil {
		return err
	}

	return nil
}

func (bs *BoardService) MakeMove(boardID string, playerID string, pieceName string, move *models.Move) error {
	board := bs.BoardMapping[boardID]
	startCell := board.Cells()[move.StartX()][move.StartY()]
	endCell := board.Cells()[move.EndX()][move.EndY()]

	currentPlayer := bs.PlayerService.PlayersMapping()[playerID]
	opponent := bs.PlayerService.PlayersMapping()[board.Player1ID()]
	if currentPlayer.Id() == board.Player1ID() {
		opponent = bs.PlayerService.PlayersMapping()[board.Player2ID()]
	}

	var currentPiece *models.Piece
	for _, piece := range currentPlayer.Pieces() {
		if piece.Name() == pieceName {
			currentPiece = &piece
			break
		}
	}

	//Remove Opponent's Piece from End Cell
	if endCell.Piece() != nil {
		piece := endCell.Piece()
		opponent.RemovePiece(piece.Id())
		endCell.SetPiece(nil)
	}

	//Move Piece
	startCell.SetPiece(nil)
	endCell.SetPiece(currentPiece)

	return nil
}

func (bs *BoardService) CheckGameStatus(boardID string, playerID string) (string, *models.Player) {
	board := bs.BoardMapping[boardID]

	currentPlayer := bs.PlayerService.PlayersMapping()[playerID]
	opponent := bs.PlayerService.PlayersMapping()[board.Player1ID()]
	if currentPlayer.Id() == board.Player1ID() {
		opponent = bs.PlayerService.PlayersMapping()[board.Player2ID()]
	}

	gameFinished, winner := bs.isCheckMate(board, currentPlayer, opponent)

	if gameFinished {
		return Finished, winner
	} else {
		return InProgress, nil
	}

}
func (bs *BoardService) isCheckMate(board *models.Board, currentPlayer *models.Player, opponent *models.Player) (bool, *models.Player) {
	var king models.Piece
	for _, piece := range opponent.Pieces() {
		if piece.Name() == models.King {
			king = piece
		}
	}
	var kingLocation models.Cell
	for _, row := range board.Cells() {
		for _, cell := range row {
			if cell.Piece() != nil {
				if cell.Piece().Id() == king.Id() {
					kingLocation = cell
				}
			}
		}
	}

	if bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) || bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) ||
		bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) || bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) ||
		bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) || bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) ||
		bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) || bs.isKingSafe(board, currentPlayer, kingLocation.X(), kingLocation.Y()) {
		return true, nil
	} else {
		return false, currentPlayer
	}

}

func (bs *BoardService) isKingSafe(board *models.Board, opponent *models.Player, kingX int, kingY int) bool {
	if kingX >= 0 && kingX < len(board.Cells()) && kingY >= 0 && kingY < len(board.Cells()) &&
		board.Cells()[kingX][kingY].Piece() == nil {

		for _, row := range board.Cells() {
			for _, cell := range row {
				if cell.X() == kingX && cell.Y() == kingY {
					continue
				}
				if cell.Piece() == nil {
					continue
				}
				pieceOnCell := cell.Piece().Name()

				err := bs.ValidateMove(board.Id(), opponent.Id(), pieceOnCell, models.NewMove(cell.X(), cell.Y(), kingX, kingY))
				if err == nil {
					return false
				}
			}
		}
	}
	return false
}

func (bs *BoardService) isMoveValid(board *models.Board, move *models.Move) error {
	cells := board.Cells()
	//validate start pos
	if move.StartX() < 0 || move.StartX() >= len(cells) || move.StartY() < 0 || move.StartY() >= len(cells) {
		return errors.New("invalid start pos")
	}

	//validate end pos
	if move.EndX() < 0 || move.EndX() >= len(cells) || move.EndY() < 0 || move.EndY() >= len(cells) {
		return errors.New("invalid end pos")
	}

	return nil
}

func checkIfCorrectPiecesPresentInMoveCells(player *models.Player, opponent *models.Player, startCell models.Cell, endCell models.Cell, pieceName string) error {
	currentPlayerPieces := player.Pieces()
	opponentPlayerPieces := opponent.Pieces()

	pieceAtStartCell := startCell.Piece().Id()
	currentPlayerPiecePresent := false
	for _, piece := range currentPlayerPieces {
		if piece.Id() == pieceAtStartCell {
			currentPlayerPiecePresent = true
			break
		}
	}
	if !currentPlayerPiecePresent {
		return errors.New("start cell don't contains current player's " + pieceName)
	}

	if endCell.Piece() != nil {
		pieceAtEndCell := endCell.Piece().Id()
		opponentPlayerPiecePresent := false
		for _, piece := range opponentPlayerPieces {
			if piece.Id() == pieceAtEndCell {
				opponentPlayerPiecePresent = true
				break
			}
		}
		if !opponentPlayerPiecePresent {
			return errors.New("end cell contains current player's " + pieceName)
		}
	}
	return nil
}
