package mongo

import (
	"apiports/internal/domain/port"
	"apiports/pkg/coordinates"
	"fmt"
)

func portBSON(p *port.Port) *BsonPort {
	doc := BsonPort{
		ID:          *p.ID,
		Name:        p.Name,
		City:        p.City,
		Country:     p.Country,
		Alias:       nil,
		Regions:     nil,
		Coordinates: nil,
		Province:    p.Province,
		Timezone:    p.Timezone,
		Unlocs:      p.Unlocs,
		Code:        p.Code,
	}

	if p.Coordinates != nil {
		doc.Coordinates = map[string]float64{
			"latitude":  p.Coordinates.Lat(),
			"longitude": p.Coordinates.Long(),
		}
	}

	if p.Alias != nil && len(p.Alias) > 0 {
		doc.Alias = make([]string, len(p.Alias))
		for i, v := range p.Alias {
			doc.Alias[i] = fmt.Sprint(v)
		}
	}

	if p.Regions != nil && len(p.Regions) > 0 {
		doc.Regions = make([]string, len(p.Regions))
		for i, v := range p.Regions {
			doc.Regions[i] = fmt.Sprint(v)
		}
	}

	return &doc
}

// ToDomainPort is a method on bson entity able to parse it back to a domain model
func (bp *BsonPort) ToDomainPort() *port.Port {
	p := port.Port{}

	p.ID = &bp.ID
	p.Name = bp.Name

	p.Regions = []interface{}{}
	for _, r := range bp.Regions {
		p.Regions = append(p.Regions, r)
	}

	p.Alias = []interface{}{}
	for _, a := range bp.Alias {
		p.Alias = append(p.Alias, a)
	}

	p.Coordinates, _ = coordinates.New([]float64{bp.Coordinates["latitude"], bp.Coordinates["longitude"]}) // skip error
	p.Code = bp.Code
	p.Unlocs = bp.Unlocs
	p.Timezone = bp.Timezone
	p.Province = bp.Province
	p.Country = bp.Country
	p.City = bp.City

	return &p
}
