package iterator

import (
	"encoding/json"
	"fmt"
	"io"
)

// Iterator is an interface for iterating through large JSON
type Iterator interface {
	NextObject() ([]byte, error)
	More() bool
}

// JSONMap implements Iterator interface
type JSONMap struct {
	decoder *json.Decoder
}

// New creates new JSONMap
func New(reader io.Reader) (*JSONMap, error) {

	// create decoder
	dec := json.NewDecoder(reader)

	// create next JSON token
	t, err := dec.Token()
	if err != nil {
		return nil, fmt.Errorf("creating next JSON token, iterator error: %w", err)
	}

	// validate delimiter
	if d, ok := t.(json.Delim); !ok || d.String() != "{" {
		return nil, fmt.Errorf("input is not an JSON map")
	}

	return &JSONMap{dec}, nil
}

// NextObject uses decoder to read and decode JSON values from an input stream
func (it *JSONMap) NextObject() ([]byte, error) {

	// create next JSON token
	tok, err := it.decoder.Token()
	if err != nil {
		return nil, fmt.Errorf("getting JSON token: %w", err)
	}

	// validate JSON map
	if val, ok := tok.(string); !ok || val == "" {
		return nil, fmt.Errorf("input is not a JSON map")
	}

	// decode raw JSON message into structure
	var raw json.RawMessage
	err = it.decoder.Decode(&raw)

	return raw, err
}

// More checks if there is more data to get
func (it *JSONMap) More() bool {
	return it.decoder.More()
}
