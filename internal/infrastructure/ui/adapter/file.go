package adapter

import (
	"apiports/internal/application/interfaces"
	"apiports/internal/domain/port"
	"apiports/internal/infrastructure/container"
	"apiports/pkg/iterator"
	"context"
	"encoding/json"
	"log"
	"os"
)

const filePath string = "assets/ports.json"

// FileImport triggers the insert process from a given raw JSON file
func FileImport(ctx context.Context, ctn *container.Container) {

	log.Println("start import")

	// open file with input data
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// create JSON iterator
	jsonStream, err := iterator.New(f)

	// while there is more bytes to read
	for jsonStream.More() {

		// read bytes representing a full JSON object
		obj, err := jsonStream.NextObject()
		if err != nil {
			log.Fatal("bad json data")
		}

		// unmarshall given object to a domain entity
		var p *port.Port
		if err := json.Unmarshal(obj, &p); err != nil {
			log.Fatal("cannot unmarshal to entity")
		}

		p.ID = &p.Unlocs[0] // quick fix as I do not iterate through a json map

		// add new port
		if err := ctn.PortService.Create(ctx, &interfaces.PortCreateDTO{Port: p}); err != nil {
			log.Fatal(err)
		}
	}

	log.Println("done import")
}
