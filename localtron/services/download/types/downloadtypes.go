/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadtypes

type ErrorResponse struct {
	Error string `json:"error"`
}

type DownloadStatus string

var (
	DownloadStatusInProgress DownloadStatus = "inProgress"
	DownloadStatusPaused     DownloadStatus = "paused"
	DownloadStatusCompleted  DownloadStatus = "completed"
	DownloadStatusErrored    DownloadStatus = "errored"
)

/* Download is the backend type for downloads */
type Download struct {
	URL            string         `json:"url"`
	FilePath       string         `json:"filePath"`
	DownloadedSize int64          `json:"downloadedSize"`
	TotalSize      int64          `json:"totalSize"`
	Status         DownloadStatus `json:"status"`
}

/* DownloadDetails is sent to the frontend */
type DownloadDetails struct {
	Id              string   `json:"id"`
	URL             string   `json:"url"`
	FileName        string   `json:"fileName"`
	Dir             *string  `json:"dir,omitempty"`
	Progress        *float64 `json:"progress,omitempty"`
	DownloadedBytes int64    `json:"downloadedBytes"`
	FullFileSize    *int64   `json:"fullFileSize"`
	Status          string   `json:"status"`
	FilePath        *string  `json:"filePath,omitempty"`
	Paused          *bool    `json:"paused,omitempty"`
	Cancelled       *bool    `json:"cancelled,omitempty"`
	Error           *string  `json:"error,omitempty"`
}

type OnFileDownloadStatus struct {
	AllDownloads []DownloadDetails `json:"allDownloads"`
}

type DownloadRequest struct {
	URL        string `json:"url"`
	FolderPath string `json:"folderPath,omitempty"`
	// FileName   *string `json:"fileName,omitempty"`
}

type GetDownloadRequest struct {
	Url string `json:"url"`
}

type GetDownloadResponse struct {
	Exists   bool             `json:"exists"`
	Download *DownloadDetails `json:"download"`
}

type DownloadResponse struct{}

type DownloadsRequest struct{}

type DownloadsResponse struct {
	Downloads []DownloadDetails `json:"downloads"`
}

//
// Events
//

const EventDownloadStatusChangeName = "downloadStatusChange"

type EventDownloadStatusChange struct {
}

func (e EventDownloadStatusChange) Name() string {
	return EventDownloadStatusChangeName
}
