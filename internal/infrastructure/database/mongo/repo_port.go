package mongo

import (
	"apiports/internal/domain/port"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName string = "default" // that should go to a config package and a different database should be provided
const collectionPorts string = "ports"

type portRepository struct {
	client *mongo.Client
}

var _ port.Repository = (*portRepository)(nil)

// NewPortRepository is a constructor for this MySQL repository
func NewPortRepository(client *mongo.Client) *portRepository {
	return &portRepository{client: client}
}

// OneByID fetches a single Port by its ID
func (r *portRepository) OneByID(ctx context.Context, id string) (*port.Port, error) {

	var bp BsonPort

	filters := bson.M{"id": id}

	if err := r.client.Database(dbName).Collection(collectionPorts).FindOne(ctx, filters).Decode(&bp); err != nil {
		return nil, err
	}

	// convert to non-DB related object
	return bp.ToDomainPort(), nil
}

// Create adds a new record in MySQL database
func (r *portRepository) Create(ctx context.Context, p *port.Port) error {

	return r.upsert(ctx, p)
}

// Update tries to update existing record or returns error if failed to find
func (r *portRepository) Update(ctx context.Context, p *port.Port) error {

	return r.upsert(ctx, p)
}

// just a shortcut - a non-production solution
func (r *portRepository) upsert(ctx context.Context, p *port.Port) error {

	bp := portBSON(p)

	filters := bson.M{"id": bp.ID}
	fields := bson.M{"$set": bp}
	opts := options.Update().SetUpsert(true)

	_, err := r.client.Database(dbName).Collection(collectionPorts).UpdateOne(ctx, filters, fields, opts)

	return err
}
