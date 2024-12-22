package models

type Player struct {
	id     string
	name   string
	color  string
	pieces []Piece
}

func NewPlayer(name string, color string) *Player {
	return &Player{name: name, color: color}
}

func (p *Player) AddPiece(piece Piece) {
	p.pieces = append(p.pieces, piece)
}

func (p *Player) RemovePiece(pieceID string) {
	var result []Piece
	for _, piece := range p.Pieces() {
		if piece.Id() != pieceID {
			result = append(result, piece)
		}
	}
	p.pieces = result
}

func (p *Player) Id() string {
	return p.id
}

func (p *Player) Name() string {
	return p.name
}

func (p *Player) Color() string {
	return p.color
}

func (p *Player) Pieces() []Piece {
	return p.pieces
}
