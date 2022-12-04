package mongo

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson"

	"apiports/internal/domain/port"
)

func portBSON(p *port.Port) *BsonPort {
	doc := BsonPort{
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
		doc.Coordinates = &bson.M{
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
