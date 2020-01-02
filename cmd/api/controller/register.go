package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	ctlmodel "github.com/superhero-register/cmd/api/model"
	"github.com/superhero-register/internal/producer/model"
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

	t := time.Now().UTC()

	// Publish superhero on Kafka topic to be stored in DB and Elasticsearch.
	err = ctl.Producer.StoreSuperhero(
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
