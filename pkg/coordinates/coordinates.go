package coordinates

import (
	"encoding/json"
	"fmt"
)

type Coordinates struct {
	latitude  float64
	longitude float64
}

var ErrInvalidInput = fmt.Errorf("excatly two float64 numbers required")

// New creates a new Coordinates object
func New(coordinates []float64) (*Coordinates, error) {

	if len(coordinates) != 2 {
		return nil, ErrInvalidInput
	}

	return &Coordinates{
		latitude:  coordinates[0],
		longitude: coordinates[1],
	}, nil
}

// UnmarshalJSON overrides deserialization process to a proper Value Object
func (c *Coordinates) UnmarshalJSON(data []byte) error {
	var cords []float64

	if err := json.Unmarshal(data, &cords); err != nil {
		return err
	}

	c.latitude = cords[0]
	c.longitude = cords[1]

	return nil
}

func (c *Coordinates) Lat() float64 {
	return c.latitude
}

func (c *Coordinates) Long() float64 {
	return c.longitude
}
