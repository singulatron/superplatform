/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3) for personal, non-commercial use.
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package firehosetypes

// Event is an interface that all events must implement
type Event interface {
	Name() string
}

// This is the event that is streamed to the frontend
type FrontendEvent struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}
