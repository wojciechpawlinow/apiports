package httpserver

import (
	"apiports/internal/application/interfaces"
	"apiports/internal/infrastructure/container"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

// ErrPortNotFound is an error when you can't find a port
var ErrPortNotFound = fmt.Errorf("port not found")

// ErrInvalidName is an error when request param is invalid
var ErrInvalidName = fmt.Errorf("port not found")

// New creates and runs a simple HTTP server
func New(ctn *container.Container) {

	e := echo.New()

	e.GET("/port/:id", func(c echo.Context) error {
		id := c.Param("id")

		if id == "" { // trivial in-place validation
			return echo.NewHTTPError(http.StatusBadRequest, ErrInvalidName.Error())
		}

		// find port in DB
		port, err := ctn.PortService.OneByID(c.Request().Context(), &interfaces.PortOneByIDDTO{ID: id})
		if port == nil {
			return echo.NewHTTPError(http.StatusNotFound, ErrPortNotFound.Error())
		}
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, fmt.Errorf("%s: %w", http.StatusText(http.StatusInternalServerError), err).Error())
		}

		return c.JSON(http.StatusOK, port)
	})

	e.Logger.Fatal(e.Start(":8080"))
}
