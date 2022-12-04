package mongo

import (
	"go.mongodb.org/mongo-driver/bson"
)

// BsonPort is a BSON representation of Port entity
type BsonPort struct {
	Name        *string  `bson:"name,omitempty"`
	City        *string  `bson:"city,omitempty"`
	Country     *string  `bson:"country,omitempty"`
	Alias       []string `bson:"alias,omitempty"`
	Regions     []string `bson:"regions,omitempty"`
	Coordinates *bson.M  `bson:"coordinates,omitempty"`
	Province    *string  `bson:"province,omitempty"`
	Timezone    *string  `bson:"timezone,omitempty"`
	Unlocs      []string `bson:"unlocs,omitempty"`
	Code        *string  `bson:"code,omitempty"`
}
