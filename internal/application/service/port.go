package service

import (
	"context"

	"apiports/internal/application/interfaces"
	"apiports/internal/domain/port"
)

// the same name of the file corresponds with the domain package it handles

type portService struct {
	PortRepository port.Repository
}

var _ interfaces.PortService = (*portService)(nil)

// NewPortService is the constructor for interfaces.PortService implementation
func NewPortService(repo port.Repository) *portService {
	return &portService{PortRepository: repo}
}

// Create takes the input data from UI and passes to the persistence layer
// in order to create a new Port
func (s *portService) Create(ctx context.Context, dto *interfaces.PortCreateDTO) error {

	return s.PortRepository.Create(ctx, dto.Port)
}

// Update takes the input data from UI and passes to the persistence layer
// in order to update an existing Port
func (s *portService) Update(ctx context.Context, dto *interfaces.PortUpdateDTO) error {

	return s.PortRepository.Update(ctx, dto.Port)
}
