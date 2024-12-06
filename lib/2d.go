package lib

type Point struct {
	X, Y int
}

func (p *Point) Scale(factor int) {
	p.X *= factor
	p.Y *= factor
}

func (p *Point) MoveRight(d int) {
	p.X += d
}

func (p *Point) MoveLeft(d int) {
	p.X -= d
}

func (p *Point) MoveUp(d int) {
	p.Y -= d
}

func (p *Point) MoveDown(d int) {
	p.Y += d
}

func (p *Point) IsNeighborDiagonal(p2 *Point, distance int) bool {
	dx := Abs(p.X - p2.X)
	dy := Abs(p.Y - p2.Y)
	return dx+dy == distance || (dx == distance && dy == distance)
}
