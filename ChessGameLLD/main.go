package main

import (
	"ChessGameLLD/models"
	"ChessGameLLD/services"
	"fmt"
)

func main() {
	playerService := services.NewPlayerService()
	boardService := services.NewBoardService(playerService)

	chessBoard := boardService.CreateBoard(3)

	playerA := playerService.NewPlayer("PlayerA", "white")
	playerB := playerService.NewPlayer("PlayerB", "black")

	pawnA := models.NewPiece(models.Pawn)
	err := playerService.AddPieceToPlayer(playerA, *pawnA)
	if err != nil {
		return
	}
	kingA := models.NewPiece(models.King)
	err = playerService.AddPieceToPlayer(playerA, *kingA)
	if err != nil {
		return
	}

	pawnB := models.NewPiece(models.Pawn)
	err = playerService.AddPieceToPlayer(playerB, *pawnB)
	if err != nil {
		return
	}
	kingB := models.NewPiece(models.King)
	err = playerService.AddPieceToPlayer(playerB, *kingB)
	if err != nil {
		return
	}

	err = boardService.PlacePieceOnBoard(chessBoard, pawnA, 0, 0)
	if err != nil {
		return
	}
	err = boardService.PlacePieceOnBoard(chessBoard, kingA, 0, 1)
	if err != nil {
		return
	}
	err = boardService.PlacePieceOnBoard(chessBoard, pawnB, 2, 0)
	if err != nil {
		return
	}
	err = boardService.PlacePieceOnBoard(chessBoard, kingB, 2, 2)
	if err != nil {
		return
	}

	status, _ := boardService.CheckGameStatus(chessBoard, playerA)
	fmt.Println(status)

}
