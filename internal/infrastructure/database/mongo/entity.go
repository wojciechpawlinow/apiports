package mongo

// BsonPort is a BSON representation of Port entity
type BsonPort struct {
	ID          string             `bson:"id"` // not the _id from Mongo, that one is transparent in this case
	Name        *string            `bson:"name,omitempty"`
	City        *string            `bson:"city,omitempty"`
	Country     *string            `bson:"country,omitempty"`
	Alias       []string           `bson:"alias,omitempty"`
	Regions     []string           `bson:"regions,omitempty"`
	Coordinates map[string]float64 `bson:"coordinates,omitempty"`
	Province    *string            `bson:"province,omitempty"`
	Timezone    *string            `bson:"timezone,omitempty"`
	Unlocs      []string           `bson:"unlocs,omitempty"`
	Code        *string            `bson:"code,omitempty"`
}
