/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */

package deploy_svc

type Command struct {
	Action      string  // e.g., "START", "KILL", "SCALE"
	ServiceSlug string  // The service affected
	Node        *string // Optional: Node where the command applies (if applicable)
	InstanceId  *string // Optional: Service instance ID (if applicable)
}
