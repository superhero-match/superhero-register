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

	"github.com/gin-gonic/gin"

	cm "github.com/superhero-match/superhero-register/internal/cache/model"
)

func (ctrl *Controller) TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenAuth, err := ctrl.Service.ExtractTokenMetadata(c.Request)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":     http.StatusUnauthorized,
				"registered": false,
			})
			c.Abort()

			return
		}

		_, err = ctrl.Service.FetchAuth(&cm.AccessDetails{
			AccessUuid: tokenAuth.AccessUuid,
			UserID:     tokenAuth.UserID,
		})
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":     http.StatusUnauthorized,
				"registered": false,
			})
			c.Abort()

			return
		}

		c.Next()
	}
}
