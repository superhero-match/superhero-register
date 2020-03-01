/*
  Copyright (C) 2019 - 2020 MWSOFT
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

type Superhero struct {
	ID                    string  `json:"id"`
	Email                 string  `json:"email"`
	Name                  string  `json:"name"`
	SuperheroName         string  `json:"superheroName"`
	MainProfilePicURL     string  `json:"mainProfilePicUrl"`
	Gender                int     `json:"gender"`
	LookingForGender      int     `json:"lookingForGender"`
	Age                   int     `json:"age"`
	LookingForAgeMin      int     `json:"lookingForAgeMin"`
	LookingForAgeMax      int     `json:"lookingForAgeMax"`
	LookingForDistanceMax int     `json:"lookingForDistanceMax"`
	DistanceUnit          string  `json:"distanceUnit"`
	Lat                   float64 `json:"lat"`
	Lon                   float64 `json:"lon"`
	Birthday              string  `json:"birthday"`
	Country               string  `json:"country"`
	City                  string  `json:"city"`
	SuperPower            string  `json:"superpower"`
	AccountType           string  `json:"accountType"`
	FirebaseToken         string  `json:"firebaseToken"`
}
