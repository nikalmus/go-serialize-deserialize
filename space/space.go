package space

import "github.com/nikalmus/go-serialize-deserialize/point"

type Space struct {
	Points []*point.Point
}

func InitSpace() *Space {
	return &Space{Points: []*point.Point{point.Origin()}}
}

func (s *Space) AddPoint(p *point.Point) {
	s.Points = append(s.Points, p)
}
