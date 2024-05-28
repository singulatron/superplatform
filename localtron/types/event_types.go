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
package types

type SelectFolderRequest struct{}

type OnFolderSelect struct {
	Location string `json:"location"`
}

type OnModelCheck struct {
	SelectedExists *bool `json:"selectedExists,omitempty"`
}

type OnModelReady struct {
	ModelReady bool `json:"modelReady"`
}

type OnOSInfo struct {
	Platform string `json:"platform"`
	Distro   string `json:"distro"`
	Release  string `json:"release"`
	Arch     string `json:"arch"`
	Hostname string `json:"hostname"`
}

type GraphicsInfoRequest struct{}

type ControllerInfo struct {
	Model       string `json:"model"`
	Vendor      string `json:"vendor"`
	Vram        int    `json:"vram"`
	VramDynamic bool   `json:"vramDynamic"`
}

type OnGraphicsInfo struct {
	Error       *string          `json:"error,omitempty"`
	Controllers []ControllerInfo `json:"controllers,omitempty"`
}

type OnSystemLanguage struct {
	SystemLanguage string `json:"systemLanguage"`
}

type DockerImagePullStatus struct {
	Status          string      `json:"status"`
	ProgressPercent *int        `json:"progressPercent,omitempty"`
	ImageName       *string     `json:"imageName,omitempty"`
	Error           interface{} `json:"error,omitempty"`
}
