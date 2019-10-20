package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/superhero-register/internal/config"
	"github.com/superhero-register/internal/producer"
)

const (
	timeFormat = "2006-01-02T15:04:05"
)

// Controller holds the controller data.
type Controller struct {
	Producer *producer.Producer
}

// NewController returns new controller.
func NewController(cfg *config.Config) (ctrl *Controller, err error) {
	return &Controller{
		Producer: producer.NewProducer(cfg),
	}, nil
}

// RegisterRoutes registers all the superheroville_municipality API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/superhero_register")

	// sr.Use(c.Authorize)

	sr.POST("/register", ctl.RegisterSuperhero)

	return router
}
