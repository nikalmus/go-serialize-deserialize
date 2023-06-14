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

func (p *Point) Serialize() ([]byte, error) {
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

func (p *Point) Deserialize(data []byte) error {
	buf := bytes.NewReader(data)

	err := binary.Read(buf, binary.LittleEndian, &p.X)
	if err != nil {
		return err
	}

	err = binary.Read(buf, binary.LittleEndian, &p.Y)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	p := Point{0, 1}

	serialized, err := p.Serialize()
	if err != nil {
		// Handle error
	}

	fmt.Printf("Serialized: %v\n", serialized)

	deserialized := Point{}
	err = deserialized.Deserialize(serialized)
	if err != nil {
		// Handle error
	}

	fmt.Printf("Deserialized: %+v\n", deserialized)
}
