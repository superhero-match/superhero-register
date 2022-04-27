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

package model

import (
	"fmt"

	validator "gopkg.in/joeybloggs/go-validate-yourself.v2"
)

type Superhero struct {
	ID                    string  `json:"id" validate:"required"`
	Email                 string  `json:"email" validate:"required"`
	Name                  string  `json:"name" validate:"required"`
	SuperheroName         string  `json:"superheroName" validate:"required"`
	MainProfilePicURL     string  `json:"mainProfilePicUrl" validate:"required"`
	Gender                int     `json:"gender" validate:"required"`
	LookingForGender      int     `json:"lookingForGender" validate:"required"`
	Age                   int     `json:"age" validate:"required"`
	LookingForAgeMin      int     `json:"lookingForAgeMin" validate:"required"`
	LookingForAgeMax      int     `json:"lookingForAgeMax" validate:"required"`
	LookingForDistanceMax int     `json:"lookingForDistanceMax" validate:"required"`
	DistanceUnit          string  `json:"distanceUnit" validate:"required"`
	Lat                   float64 `json:"lat" validate:"required"`
	Lon                   float64 `json:"lon" validate:"required"`
	Birthday              string  `json:"birthday" validate:"required"`
	Country               string  `json:"country" validate:"required"`
	City                  string  `json:"city" validate:"required"`
	SuperPower            string  `json:"superpower" validate:"required"`
	AccountType           string  `json:"accountType" validate:"required"`
	FirebaseToken         string  `json:"firebaseToken" validate:"required"`
	IsDeleted             bool    `json:"isDeleted" validate:"omitempty"`
	DeletedAt             string  `json:"deletedAt" validate:"omitempty"`
	IsBlocked             bool    `json:"isBlocked" validate:"omitempty"`
	BlockedAt             string  `json:"blockedAt" validate:"omitempty"`
	UpdatedAt             string  `json:"updatedAt" validate:"omitempty"`
	CreatedAt             string  `json:"createdAt" validate:"omitempty"`
}

// Validate validates that all fields are present.
func (s Superhero) Validate() error {
	errs := validator.ValidateStruct(s)
	if errs != nil {
		return fmt.Errorf("error validating Superhero: %v", errs.Errors)
	}

	return nil
}
