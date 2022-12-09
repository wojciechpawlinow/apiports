package container

import (
	"fmt"
	"log"
	
	"apiports/internal/application/interfaces"
	"apiports/internal/application/service"
	"apiports/internal/infrastructure/database/mongo"
)

// Container is a special object that keeps all dependencies)
type Container struct {

	// services
	PortService interfaces.PortService

	// ... more
}

// New creates all dependencies and should also return error in case of any failure at this stage
func New() *Container {

	// handle errors gracefully - building container and dependencies
	// can also contain some Ping() methods in order to check readiness
	conn, err := mongo.Connect()
	if err != nil || conn == nil {
		log.Fatal(fmt.Errorf("dependency failed: %w", err)) // I never use panic, but here I just want to crash the app if cant connect to mongo
	}
	// create repository
	repoPort := mongo.NewPortRepository(conn)

	// create application service
	servicePort := service.NewPortService(repoPort)

	return &Container{
		PortService: servicePort,
	}
}
