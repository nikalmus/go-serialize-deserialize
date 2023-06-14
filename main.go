package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Point struct {
	X int32
	Y int32
}

func SerializePoint(p Point) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.LittleEndian, p.X)
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.LittleEndian, p.Y)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func DeserializePoint(data []byte) (Point, error) {
	buf := bytes.NewReader(data)

	var p Point

	err := binary.Read(buf, binary.LittleEndian, &p.X)
	if err != nil {
		return Point{}, err
	}

	
	err = binary.Read(buf, binary.LittleEndian, &p.Y)
	if err != nil {
		return Point{}, err
	}

	return p, nil
}

func main() {
	p := Point{0, 1}

	serialized, err := SerializePoint(p)
	if err != nil {
		// Handle error
	}

	fmt.Printf("Serialized : %v\n", serialized)

	deserialized, err := DeserializePoint(serialized)
	if err != nil {
		// Handle error
	}

	fmt.Printf("Deserialized : %v\n", deserialized)
}