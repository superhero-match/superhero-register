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
package service

import (
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/superhero-match/superhero-register/internal/cache"
	chm "github.com/superhero-match/superhero-register/internal/cache/model"
	"github.com/superhero-match/superhero-register/internal/config"
	"github.com/superhero-match/superhero-register/internal/producer"
	"github.com/superhero-match/superhero-register/internal/producer/model"
)

// Service interface defines service methods.
type Service interface {
	FetchAuth(authD *chm.AccessDetails) (string, error)
	ExtractToken(r *http.Request) string
	VerifyToken(r *http.Request) (*jwt.Token, error)
	ExtractTokenMetadata(r *http.Request) (*chm.AccessDetails, error)
	Close() error
	StoreSuperhero(s model.Superhero) error
}

// service holds all the different services that are used when handling request.
type service struct {
	Producer     producer.Producer
	Cache        cache.Cache
	AccessSecret string
}

// NewService creates value of type Service.
func NewService(cfg *config.Config) (Service, error) {
	c, err := cache.NewCache(cfg)
	if err != nil {
		return nil, err
	}

	return &service{
		Producer:     producer.NewProducer(cfg),
		Cache:        c,
		AccessSecret: cfg.JWT.AccessTokenSecret,
	}, nil
}
