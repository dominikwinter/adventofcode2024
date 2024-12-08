package lib

type Node struct {
	X, Y int
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
	n.X += v.X
	n.Y += v.Y

	return n
}

func (n *Node) Backward(v *Node) *Node {
	n.X -= v.X
	n.Y -= v.Y

	return n
}

func (n *Node) SetXY(n2 *Node) *Node {
	n.X = n2.X
	n.Y = n2.Y

	return n
}

func (n *Node) IsInBounds(m *[][]string) bool {
	return n.Y >= 0 && n.Y < len(*m) && n.X >= 0 && n.X < len((*m)[n.Y])
}
