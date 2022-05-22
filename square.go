package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (receiver Square) End() Point {
	intA := int(receiver.a)
	return Point{x: receiver.start.x + intA, y: receiver.start.y + intA}

}

func (receiver Square) Area() uint {
	return receiver.a * receiver.a
}

func (receiver Square) Perimeter() uint {
	return receiver.a * 4
}
