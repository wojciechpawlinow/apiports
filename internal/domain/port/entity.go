package port

import "apiports/pkg/coordinates"

// Port is the domain entity
// The assumption is there are no nullable fields (that's why there are no pointers)
// Moreover, City, Country, Regions, Coordinates, Timezone and Province could be
// potentially merged into a separate Location value object
type Port struct {
	ID          *string                  `json:"id"`
	Name        *string                  `json:"name"`
	City        *string                  `json:"city"`
	Country     *string                  `json:"country"`
	Alias       []interface{}            `json:"alias"`       // type not recognised from the JSON file
	Regions     []interface{}            `json:"regions"`     // type not recognised from the JSON file
	Coordinates *coordinates.Coordinates `json:"coordinates"` // value object created
	Province    *string                  `json:"province"`
	Timezone    *string                  `json:"timezone"`
	Unlocs      []string                 `json:"unlocs"`
	Code        *string                  `json:"code"`
}
