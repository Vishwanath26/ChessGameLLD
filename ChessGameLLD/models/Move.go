package models

import "github.com/google/uuid"

type Move struct {
	id     string
	startX int
	startY int
	endX   int
	endY   int
}

func (m Move) Id() string {
	return m.id
}

func (m Move) StartX() int {
	return m.startX
}

func (m Move) StartY() int {
	return m.startY
}

func (m Move) EndX() int {
	return m.endX
}

func (m Move) EndY() int {
	return m.endY
}

func NewMove(startX int, startY int, endX int, endY int) *Move {
	return &Move{
		id:     uuid.NewString(),
		startX: startX,
		startY: startY,
		endX:   endX,
		endY:   endY,
	}
}
