package main

import (
	"fmt"
	"github.com/nikalmus/go-serialize-deserialize/point"
	"github.com/nikalmus/go-serialize-deserialize/space"
	//"github.com/dgraph-io/badger/v3"
)

func main() {
	// p := point.Point{X: 0, Y: 1}

	// serialized, err := p.Serialize()
	// if err != nil {
	// 	// Handle error
	// }

	// fmt.Printf("Serialized: %v\n", serialized)

	// deserialized := point.Point{}
	// err = deserialized.Deserialize(serialized)
	// if err != nil {
	// 	// Handle error
	// }

	// fmt.Printf("Deserialized: %+v\n", deserialized)
	p1 := &point.Point{X: 1, Y: 2}
	p2 := &point.Point{X: 3, Y: 4}

	s := space.InitSpace()
	s.AddPoint(p1)
	s.AddPoint(p2)

	for _, p := range s.Points {
		fmt.Printf("Point: %+v\n", *p)
	}
}

