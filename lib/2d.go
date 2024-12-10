package lib

type Node struct {
	Y, X int
}

func (n *Node) Copy() *Node {
	return &Node{n.X, n.Y}
}

func (n *Node) Scale(factor int) *Node {
	n.X *= factor
	n.Y *= factor

	return n
}

func (n *Node) MoveRight(d int) *Node {
	n.X += d

	return n
}

func (n *Node) MoveLeft(d int) *Node {
	n.X -= d

	return n
}

func (n *Node) MoveUp(d int) *Node {
	n.Y -= d

	return n
}

func (n *Node) MoveDown(d int) *Node {
	n.Y += d

	return n
}

func (n *Node) Forward(v *Node) *Node {
	n.Y += v.Y
	n.X += v.X

	return n
}

func (n *Node) Backward(v *Node) *Node {
	n.Y -= v.Y
	n.X -= v.X

	return n
}

func (n *Node) SetXY(n2 *Node) *Node {
	n.Y = n2.Y
	n.X = n2.X

	return n
}

// y x
var Dir = [][]int{
	{-1, 0}, // up
	{1, 0},  // down
	{0, -1}, // left
	{0, 1},  // right
}

func (n *Node) GetNeighbors(m *[][]int) []*Node {
	var neighbors []*Node

	for _, d := range Dir {
		neighbor := &Node{
			Y: n.Y + d[0],
			X: n.X + d[1],
		}

		if neighbor.IsInBounds(m) {
			neighbors = append(neighbors, neighbor)
		}
	}

	return neighbors
}

func (n *Node) GetValue(m *[][]int) int {
	return (*m)[n.Y][n.X]
}

func (n *Node) IsInBounds(m *[][]int) bool {
	return n.Y >= 0 && n.Y < len(*m) && n.X >= 0 && n.X < len((*m)[n.Y])
}

// TODO this sucks
func (n *Node) IsInBoundsStr(m *[][]string) bool {
	return n.Y >= 0 && n.Y < len(*m) && n.X >= 0 && n.X < len((*m)[n.Y])
}
