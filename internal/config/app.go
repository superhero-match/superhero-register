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
package config

// App holds the configuration values for the application.
type App struct {
	Port       string `env:"APP_PORT" default:":3000"`
	CertFile   string `env:"APP_CERT_FILE" default:"./cmd/api/certificate.pem"`
	KeyFile    string `env:"APP_KEY_FILE" default:"./cmd/api/key.pem"`
	TimeFormat string `env:"APP_TIME_FORMAT" default:"2019-09-15T14:04:05"`
}