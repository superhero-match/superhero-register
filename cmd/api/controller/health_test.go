/*
  Copyright (C) 2019 - 2022 MWSOFT
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
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func MockGet(c *gin.Context) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
}

func TestController_Health(t *testing.T) {
	mockProd := mockProducer{
		registerSuperhero: mockPublishRegisterSuperhero,
	}

	mService := &mockService{
		mProducer: mockProd,
	}

	logger, err := zap.NewProduction()
	if err != nil {
		t.Fatal(err)
	}

	defer logger.Sync()

	mockController := &Controller{
		Service:    mService,
		Logger:     logger,
		TimeFormat: "2006-01-02T15:04:05",
	}

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	ctx.Request = &http.Request{
		Header: make(http.Header),
	}

	MockGet(ctx)

	mockController.Health(ctx)
	assert.EqualValues(t, http.StatusOK, w.Code)
}
