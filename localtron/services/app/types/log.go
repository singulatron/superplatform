/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package apptypes

type Log struct {
	Level    string         `json:"level"`
	Time     string         `json:"time"`
	Source   string         `json:"source"`
	Ip       string         `json:"ip"`
	ClientId string         `json:"clientId"`
	Platform string         `json:"platform"`
	Message  string         `json:"message"`
	Fields   map[string]any `json:"fields"`
}

type LogRequest struct {
	Logs []Log `json:"logs"`
}

type LoggingStatus struct {
	Enabled bool `json:"enabled"`
}

type AppServiceI interface {
	LoggingStatus() (LoggingStatus, error)
	EnableLogging() error
	DisableLogging() error
}
