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
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/superhero-match/superhero-register/internal/cache/model"
)

func (srv *service) ExtractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")

	strArr := strings.Split(bearToken, " ")

	if len(strArr) == 2 {
		return strArr[1]
	}

	return ""
}

func (srv *service) VerifyToken(r *http.Request) (*jwt.Token, error) {
	tokenString := srv.ExtractToken(r)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(srv.AccessSecret), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (srv *service) ExtractTokenMetadata(r *http.Request) (*model.AccessDetails, error) {
	token, err := srv.VerifyToken(r)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		accessUuid, ok := claims["access_uuid"].(string)
		if !ok {
			return nil, fmt.Errorf("could not fetch access uuid from claims")
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			return nil, fmt.Errorf("could not fetch user id from claims")
		}

		return &model.AccessDetails{
			AccessUuid: accessUuid,
			UserID:     userId,
		}, nil
	}

	return nil, err
}
