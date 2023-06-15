package point

import (
	"bytes"
	"encoding/binary"
)

type Point struct {
	X int32
	Y int32
}

func CreatePoint(x, y int32) *Point {
	return &Point{X: x, Y: y}
}

func Origin() *Point {
	return CreatePoint(0, 0)
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