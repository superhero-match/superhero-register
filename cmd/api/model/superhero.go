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
