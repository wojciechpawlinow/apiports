package interfaces

import (
	"apiports/internal/domain/port"
	"context"
)

// the name of this file corresponds with the service it exposes as an interface

// define DTO interfaces for accessing application layer live
// in the same space (application layer) where the actual implementation lives
// that prevents leaking the data from the top to bottom layers and so that prevents coupling as well

// PortOneByIDDTO is the API contract for getting Port
type PortOneByIDDTO struct {
	ID string
}

// PortCreateDTO is the API contract for creating Port
type PortCreateDTO struct {
	Port *port.Port
}

// PortUpdateDTO is the API contract for updating Port
type PortUpdateDTO struct {
	Port *port.Port

	// that could potentially have different data (i.e partial Port's data)
	// all depends on business requirements
}

// PortService defines the access pattern for Port domain. It allows to perform a business logic
// in a separation from any User Interface (Ports & Adapters architecture) -
// this place stands the Port, to which you can plug many adapters
type PortService interface {
	OneByID(ctx context.Context, dto *PortOneByIDDTO) (*port.Port, error)
	Create(ctx context.Context, dto *PortCreateDTO) error
	Update(ctx context.Context, dto *PortUpdateDTO) error
}
