/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package node_types

type Options struct {
	Port        int
	GpuPlatform string
	// Singulatron Server Address
	Address    string
	Az         string
	Region     string
	LLMHost    string
	VolumeName string
	ConfigPath string
	Db         string
	DbDriver   string
	DbString   string

	// DbPrefix allows us to have isolated envs for different test cases
	// but still make multiple nodes in those test cases use the same
	// shard of the db.
	DbPrefix string
}
