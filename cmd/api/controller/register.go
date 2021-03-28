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
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ctlmodel "github.com/superhero-match/superhero-register/cmd/api/model"
	"github.com/superhero-match/superhero-register/internal/producer/model"
)

const (
	minimumRegistrationAge int = 18
	timeFormat string = "2006-01-02T15:04:05"
	timeFormatShort string = "2006-01-02"
	daysInAYear int = 365
	hoursInADay float64 = 24
)

// RegisterSuperhero registers new Superhero.
func (ctl *Controller) RegisterSuperhero(c *gin.Context) {
	var s ctlmodel.Superhero

	err := c.BindJSON(&s)
	if err != nil {
		fmt.Println("BindJSON")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}

	fmt.Println()
	fmt.Printf("%+v", s)
	fmt.Println()

	fmt.Println("before s.Age < minimumRegistrationAge")
	if s.Age < minimumRegistrationAge {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}
	fmt.Println("after s.Age < minimumRegistrationAge")

	fmt.Println("before time.Parse(timeFormatShort, s.Birthday)")
	now := time.Now()
	then, err := time.Parse(timeFormatShort, s.Birthday)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}
	fmt.Println("before time.Parse(timeFormatShort, s.Birthday)")

	diff := now.Sub(then)
	days := int(diff.Hours() / hoursInADay)
	years := int(days / daysInAYear)

	fmt.Println("before if years < minimumRegistrationAge")
	if years < minimumRegistrationAge {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}
	fmt.Println("after if years < minimumRegistrationAge")

	t := time.Now().UTC()

	// Publish superhero on Kafka topic to be stored in DB and Elasticsearch.
	err = ctl.Service.Producer.StoreSuperhero(
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
	if err != nil {
		fmt.Println("err -> ctl.Service.Producer.StoreSuperhero")
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":     http.StatusInternalServerError,
			"registered": false,
		})

		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":     http.StatusOK,
		"registered": true,
	})
}
