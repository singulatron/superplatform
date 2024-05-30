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
package features

import dt "github.com/singulatron/singulatron/dapper/types"

var WslSetDefaultVersion = dt.Feature{
	ID:   "wsl-set-default-version",
	Name: "WSL Set Default Version",
	Arguments: []dt.Argument{
		{
			Name:    "wslVersion",
			Type:    dt.Int,
			Default: 2,
		},
	},
	PlatformScripts: map[dt.Platform]*dt.Scripts{
		dt.Windows: {
			Execute: &dt.Script{
				Source: `
Write-Host "Setting default WSL version to {{.wslVersion}}"
wsl --set-default-version {{.wslVersion}}
`,
				Runtime: "powershell",
			},
			Check: &dt.Script{
				Source: `
$wslDefaultVersionCheck = wsl --list --verbose | Select-String "Default"
if ($wslDefaultVersionCheck -match "v{{.wslVersion}}") {
    # Default version is already set to {{.wslVersion}}
    Write-Host "yes"
} else {
    # Default version is not set to {{.wslVersion}}
    Write-Host "no"
}`,
				Runtime: "powershell",
			},
		},
	},
	PlatformFeatures: map[dt.Platform][]any{
		dt.Windows: {
			map[string]any{
				"featureId": WslUpdated.ID,
				"args": []any{
					"2",
				},
			},
		},
	},
}
