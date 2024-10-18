/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

type CommandType string

const (
	CommandTypeStart CommandType = "START"
	CommandTypeKill  CommandType = "KILL"
	CommandTypeScale CommandType = "SCALE"
)

type Command struct {
	Action      CommandType // e.g., "START", "KILL", "SCALE"
	ServiceSlug string      // The User Svc slug of the service affected
	NodeUrl     *string     // Node address, e.g., "https://api.com:999"
	InstanceId  *string     // Instance id, e.g., "https://api.com:999/user-svc"
}
