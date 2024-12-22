package services

import (
	"ChessGameLLD/models"
	"errors"
)

type PlayerService struct {
	playersMapping map[string]*models.Player
}

func (ps *PlayerService) PlayersMapping() map[string]*models.Player {
	return ps.playersMapping
}

func NewPlayerService() IPlayerService {
	return &PlayerService{playersMapping: make(map[string]*models.Player)}
}

type IPlayerService interface {
	NewPlayer(name string, color string) string
	AddPieceToPlayer(playerID string, piece models.Piece) error
	PlayersMapping() map[string]*models.Player
}

func (ps *PlayerService) NewPlayer(name string, color string) string {
	newPlayer := models.NewPlayer(name, color)
	ps.playersMapping[newPlayer.Id()] = newPlayer
	return newPlayer.Id()
}

func (ps *PlayerService) AddPieceToPlayer(playerID string, piece models.Piece) error {
	if ps.playersMapping[playerID] == nil {
		return errors.New("invalid player id")
	}
	player := ps.playersMapping[playerID]
	player.AddPiece(piece)

	return nil
}
