package main

import (
	"os"
	"log"
	"fmt"
	"github.com/nikalmus/go-serialize-deserialize/point"
	"github.com/nikalmus/go-serialize-deserialize/space"
)

func main() {
	dbPath := "./db/pointstore" 

	// Create the directory if it doesn't exist
	err := os.MkdirAll(dbPath, 0755)
	if err != nil {
		log.Fatal(err)
	}

	s, err := space.InitSpace(dbPath)
	if err != nil {
		log.Fatal(err)
	}

	p1 := &point.Point{X: 11, Y: 21}
	p2 := &point.Point{X: 35, Y: 45}

	// Call AddPoint to save the points
	err = s.AddPoint(p1)
	if err != nil {
		log.Fatal(err)
	}

	err = s.AddPoint(p2)
	if err != nil {
		log.Fatal(err)
	}

	// Call LoadPoints to retrieve the saved points
	points, err := s.LoadPoints()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Points in the database:")

	for _, p := range points {
		fmt.Printf("X: %d, Y: %d\n", p.X, p.Y)
	}

	err = s.Close()
	if err != nil {
		log.Fatal(err)
	}
}

