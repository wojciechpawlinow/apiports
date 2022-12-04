package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"apiports/internal/application/interfaces"
	"apiports/internal/domain/port"
)

// PortRepository is a repository mock structure
type PortRepository struct {
	mock.Mock
}

var _ port.Repository = (*PortRepository)(nil)

// Create is a mock method implementing the interface
func (m *PortRepository) Create(ctx context.Context, p *port.Port) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

// Update is a mock method implementing the interface
func (m *PortRepository) Update(ctx context.Context, p *port.Port) error {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func TestPortService_Create(t *testing.T) {

	// here could be a helper func to build all dependencies once and decrease code redundancy
	t.Run("happy path", func(t *testing.T) {
		ctx := context.TODO()
		repoMock := new(PortRepository)

		testName := "test"
		testPort := port.Port{Name: &testName}

		repoMock.On("Create", mock.Anything, &testPort).Return(nil)

		srv := NewPortService(repoMock)

		err := srv.Create(ctx, &interfaces.PortCreateDTO{Port: &testPort})

		assert.Nil(t, err)
		repoMock.AssertExpectations(t)
	})

	t.Run("failed to create", func(t *testing.T) {
		ctx := context.TODO()
		repoMock := new(PortRepository)

		testName := "test"
		testPort := port.Port{Name: &testName}

		repoMock.On("Create", mock.Anything, &testPort).Return(fmt.Errorf("ugly fail"))

		srv := NewPortService(repoMock)

		err := srv.Create(ctx, &interfaces.PortCreateDTO{Port: &testPort})

		assert.Error(t, err)
		repoMock.AssertExpectations(t)
	})
}

func TestPortService_Update(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		ctx := context.TODO()
		repoMock := new(PortRepository)

		testName := "test"
		testPort := port.Port{Name: &testName}

		repoMock.On("Update", mock.Anything, &testPort).Return(nil)

		srv := NewPortService(repoMock)

		err := srv.Update(ctx, &interfaces.PortUpdateDTO{Port: &testPort})

		assert.Nil(t, err)
		repoMock.AssertExpectations(t)
	})

	t.Run("failed to update", func(t *testing.T) {
		ctx := context.TODO()
		repoMock := new(PortRepository)

		testName := "test"
		testPort := port.Port{Name: &testName}

		repoMock.On("Update", mock.Anything, &testPort).Return(fmt.Errorf("ugly fail"))

		srv := NewPortService(repoMock)

		err := srv.Update(ctx, &interfaces.PortUpdateDTO{Port: &testPort})

		assert.Error(t, err)
		repoMock.AssertExpectations(t)
	})
}
