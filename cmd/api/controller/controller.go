/*
  Copyright (C) 2019 - 2021 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/superhero-match/superhero-register/cmd/api/service"

	"github.com/superhero-match/superhero-register/internal/config"
)

const (
	timeFormat = "2006-01-02T15:04:05"
)

// Controller holds the controller data.
type Controller struct {
	Service *service.Service
}

// NewController returns new controller.
func NewController(cfg *config.Config) (ctrl *Controller, err error) {
	srv, err := service.NewService(cfg)
	if err != nil {
		return nil, err
	}

	return &Controller{
		Service: srv,
	}, nil
}

// RegisterRoutes registers all the superhero register API routes.
func (ctl *Controller) RegisterRoutes() *gin.Engine {
	router := gin.Default()

	sr := router.Group("/api/v1/superhero_register")

	// sr.Use(c.Authorize)

	sr.POST("/register", ctl.RegisterSuperhero)

	return router
}
