package space

import (
	"fmt"
	"github.com/dgraph-io/badger/v3"
	"github.com/nikalmus/go-serialize-deserialize/point"
)

type Space struct {
	Points []*point.Point
	db     *badger.DB
}

func InitSpace(dbPath string) (*Space, error) {
	opt := badger.DefaultOptions(dbPath)
	opt = opt.WithLoggingLevel(badger.WARNING)

	db, err := badger.Open(opt)

	if err != nil {
		return nil, err
	}

	s := &Space{
		Points: []*point.Point{},
		db:     db,
	}

	isEmpty, err := s.isDatabaseEmpty()
	if err != nil {
		return nil, err
	}

	if isEmpty {
		err = s.savePoint(point.Origin()) 
		if err != nil {
			return nil, err
		}
	} else {

		points, err := s.LoadPoints()
		if err != nil {
			return nil, err
		}
		s.Points = points
	}

	return s, nil
}

func (s *Space) isDatabaseEmpty() (bool, error) {
	var isEmpty bool

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		iterator := txn.NewIterator(opts)
		defer iterator.Close()

		isEmpty = !iterator.Valid()
		return nil
	})

	return isEmpty, err
}

func (s *Space) LoadPoints() ([]*point.Point, error) {
	fmt.Println("inside LoadPoints")
	points := []*point.Point{}

	err := s.db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false

		iterator := txn.NewIterator(opts)
		defer iterator.Close()

		for iterator.Seek([]byte("")); iterator.Valid(); iterator.Next() {
			item := iterator.Item()

			err := item.Value(func(val []byte) error {
				deserialized := &point.Point{}
				err := deserialized.Deserialize(val)
				if err != nil {
					return err
				}
                fmt.Println("Loaded point with X:", deserialized.X, "Y:", deserialized.Y)
				points = append(points, deserialized)

				return nil
			})
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to load points from the database: %w", err)
	}

	return points, nil
}



func (s *Space) savePoint(p *point.Point) error {
	fmt.Println("inside savePoint")
	serialized, err := p.Serialize()
	if err != nil {
		return err
	}

	fmt.Println("Saving point with X:", p.X, "Y:", p.Y, "Serialized:", serialized)
	err = s.db.Update(func(txn *badger.Txn) error {
		err := txn.Set(serialized, serialized)
		return err
	})

	if err != nil {
		return fmt.Errorf("failed to save point: %w", err)
	}

	return nil
}


func (s *Space) Close() error {
	if s.db != nil {
		err := s.db.Close()
		s.db = nil
		return err
	}
	return nil
}

func (s *Space) AddPoint(p *point.Point) error {
	exists, err := s.PointExists(p)
	if err != nil {
		return fmt.Errorf("failed to check if point exists: %w", err)
	}

	if exists {
		fmt.Printf("Skipping duplicate point with X: %d, Y: %d\n", p.X, p.Y)
		return nil // Skip the duplicate point
	}

	s.Points = append(s.Points, p)
	err = s.savePoint(p)
	if err != nil {
		return fmt.Errorf("failed to save point to the database: %w", err)
	}

	return nil
}


func (s *Space) PointExists(p *point.Point) (bool, error) {
	db := s.db

	var exists bool

	err := db.View(func(txn *badger.Txn) error {
		serialized, err := p.Serialize()
		if err != nil {
			return err
		}

		_, err = txn.Get(serialized)
		if err == nil {
			exists = true
		} else if err == badger.ErrKeyNotFound {
			exists = false
		} else {
			return err
		}

		return nil
	})

	if err != nil {
		return false, fmt.Errorf("failed to check if point exists: %w", err)
	}

	return exists, nil
}




