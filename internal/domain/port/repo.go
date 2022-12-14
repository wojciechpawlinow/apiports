package port

import "context"

// for this following task is sufficient to provide only a sing method able to either

// Repository is an interface of the Port domain's persistence layer
type Repository interface {
	OneByID(ctx context.Context, id string) (*Port, error)
	Create(ctx context.Context, p *Port) error
	Update(ctx context.Context, p *Port) error
	// could be also  UpsertOne(ctx context.Context, port *Port) error
}
