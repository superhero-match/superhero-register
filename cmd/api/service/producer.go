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
	"github.com/superhero-match/superhero-register/internal/producer/model"
)

// Close closes the connection to Kafka.
func (srv *service) Close() error {
	return srv.Producer.Close()
}

// StoreSuperhero publishes new Superhero on Kafka topic for it to be
// consumed by record keeper and stored in DB and Elasticsearch.
func (srv *service) StoreSuperhero(s model.Superhero) error {
	return srv.Producer.StoreSuperhero(s)
}
