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

package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	err := os.Setenv("TEST_CONFIG", "config.test.yml")
	if err != nil {
		t.Fatal(err)
	}

	cfg, err := New()
	if err != nil {
		t.Fatal(err)
	}

	// App configuration.
	assert.Equal(t, ":3000", cfg.App.Port, "The port should be :3000.")
	assert.Equal(t, "2006-01-02T15:04:05", cfg.App.TimeFormat, "The time format should be 2006-01-02T15:04:05.")

	// Kafka producer.
	assert.Equal(t, "localhost:9092", cfg.Producer.Brokers, "The Kafka producer brokers should be localhost:9092.")
	assert.Equal(t, "store.municipality.superhero", cfg.Producer.Topic, "The Kafka producer topic should be store.municipality.superhero.")
	assert.Equal(t, 1, cfg.Producer.BatchSize, "The Kafka producer batch size should be 1.")
	assert.Equal(t, 10, cfg.Producer.BatchTimeout, "The Kafka producer batch timeout should be 10.")
}
