/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 *
 * For commercial use, a separate license must be obtained by purchasing from The Authors.
 * For commercial licensing inquiries, please contact The Authors listed in the AUTHORS file.
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
