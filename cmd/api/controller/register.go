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
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	ctlmodel "github.com/superhero-match/superhero-register/cmd/api/model"
	"github.com/superhero-match/superhero-register/internal/producer/model"
)

const (
	minimumRegistrationAge int     = 18
	timeFormat             string  = "2006-01-02T15:04:05"
	timeFormatShort        string  = "2006-01-02"
	daysInAYear            int     = 365
	hoursInADay            float64 = 24
)

// RegisterSuperhero registers new Superhero.
func (ctl *Controller) RegisterSuperhero(c *gin.Context) {
	var s ctlmodel.Superhero

	err := c.BindJSON(&s)
	if checkError(err, c) {
		ctl.Logger.Error(
			"failed to bind request model",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	if s.Age < minimumRegistrationAge {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}

	now := time.Now()
	then, err := time.Parse(timeFormatShort, s.Birthday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}

	diff := now.Sub(then)
	days := int(diff.Hours() / hoursInADay)
	years := int(days / daysInAYear)

	if years < minimumRegistrationAge {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}

	t := time.Now().UTC()

	// Publish superhero on Kafka topic to be stored in DB and Elasticsearch.
	err = ctl.Service.StoreSuperhero(
		model.Superhero{
			ID:                    s.ID,
			Email:                 s.Email,
			Name:                  s.Name,
			SuperheroName:         s.SuperheroName,
			MainProfilePicURL:     s.MainProfilePicURL,
			Gender:                s.Gender,
			LookingForGender:      s.LookingForGender,
			Age:                   s.Age,
			LookingForAgeMin:      s.LookingForAgeMin,
			LookingForAgeMax:      s.LookingForAgeMax,
			LookingForDistanceMax: s.LookingForDistanceMax,
			DistanceUnit:          s.DistanceUnit,
			Lat:                   s.Lat,
			Lon:                   s.Lon,
			Birthday:              s.Birthday,
			Country:               s.Country,
			City:                  s.City,
			SuperPower:            s.SuperPower,
			AccountType:           s.AccountType,
			FirebaseToken:         s.FirebaseToken,
			IsDeleted:             false,
			DeletedAt:             string(""),
			IsBlocked:             false,
			BlockedAt:             string(""),
			UpdatedAt:             t.Format(timeFormat),
			CreatedAt:             t.Format(timeFormat),
		},
	)
	time.Now().UTC()
	if checkError(err, c) {
		ctl.Logger.Error(
			"failed to store superhero",
			zap.String("err", err.Error()),
			zap.String("time", time.Now().UTC().Format(ctl.TimeFormat)),
		)

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"registered": true,
	})
}

func checkError(err error, c *gin.Context) bool {
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return true
	}

	return false
}
