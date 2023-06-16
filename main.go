package main

import (
	"github.com/nikalmus/go-serialize-deserialize/point"
	"github.com/nikalmus/go-serialize-deserialize/space"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	dbPath := "./db/pointstore" 

	s, err := space.InitSpace(dbPath)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	createCmd := flag.NewFlagSet("create", flag.ExitOnError)
	loadCmd := flag.NewFlagSet("load", flag.ExitOnError)

	x := createCmd.Int("x", 0, "X coordinate")
	y := createCmd.Int("y", 0, "Y coordinate")

	if len(os.Args) < 2 {
		fmt.Println("create or load command is required")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "create":
		createCmd.Parse(os.Args[2:])
		p := point.CreatePoint(int32(*x), int32(*y))
		err := s.AddPoint(p)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Point created successfully!")
	case "load":
		loadCmd.Parse(os.Args[1:])
		points, err := s.LoadPoints()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Points in the database:")
		for _, p := range points {
			fmt.Printf("X: %d, Y: %d\n", p.X, p.Y)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}
}

