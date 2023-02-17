package main

type Geometry interface {
	Edges() int
}
type Pentagon struct{}
type Hexagon struct{}
type Decagon struct{}
type Octagon struct{}

func (p Pentagon) Edges() int {
	return 5
}
func (h Hexagon) Edges() int {
	return 6
}
func (o Octagon) Edges() int {
	return 8
}
func (d Decagon) Edges() int {
	return 10
}
func calc(geo Geometry, val int) int {
	num := geo.Edges()
	res := num * val
	return res
}
func main() {

}
