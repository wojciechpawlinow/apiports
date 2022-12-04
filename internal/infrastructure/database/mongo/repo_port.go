package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"apiports/internal/domain/port"
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

// Create adds a new record in MySQL database
func (r *portRepository) Create(ctx context.Context, p *port.Port) error {

	_, err := r.client.Database(dbName).Collection(collectionPorts).InsertOne(ctx, portBSON(p))

	return err
}

// Update tries to update existing record or returns error if failed to find
func (*portRepository) Update(_ context.Context, _ *port.Port) error {

	// similar as above but maybe with UpdateOne

	return nil
}
