package models

import "github.com/google/uuid"

type Board struct {
	id         string
	cells      [][]Cell
	lastPlayer string
	player1ID  string
	player2ID  string
}

func (b *Board) Cells() [][]Cell {
	return b.cells
}

func (b *Board) Player1ID() string {
	return b.player1ID
}

func (b *Board) SetPlayer1ID(player1ID string) {
	b.player1ID = player1ID
}

func (b *Board) Player2ID() string {
	return b.player2ID
}

func (b *Board) SetPlayer2ID(player2ID string) {
	b.player2ID = player2ID
}

func (b *Board) Id() string {
	return b.id
}

func (b *Board) LastPlayer() string {
	return b.lastPlayer
}

func (b *Board) SetLastPlayer(lastPlayer string) {
	b.lastPlayer = lastPlayer
}

func NewBoard(cells [][]Cell) *Board {
	return &Board{
		id:    uuid.NewString(),
		cells: cells,
	}
}
