package port

import "apiports/pkg/coordinates"

// Port is the domain entity
// The assumption is there are no nullable fields (that's why there are no pointers)
// Moreover, City, Country, Regions, Coordinates, Timezone and Province could be
// potentially merged into a separate Location value object
type Port struct {
	Name        *string                  `json:"name,omitempty"`
	City        *string                  `json:"city,omitempty"`
	Country     *string                  `json:"country,omitempty"`
	Alias       []interface{}            `json:"alias,omitempty"`       // type not recognised from the JSON file
	Regions     []interface{}            `json:"regions,omitempty"`     // type not recognised from the JSON file
	Coordinates *coordinates.Coordinates `json:"coordinates,omitempty"` // value object created
	Province    *string                  `json:"province,omitempty"`
	Timezone    *string                  `json:"timezone,omitempty"`
	Unlocs      []string                 `json:"unlocs,omitempty"`
	Code        *string                  `json:"code,omitempty"`
}
