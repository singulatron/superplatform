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
package downloadservice

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/singulatron/singulatron/localtron/logger"
	types "github.com/singulatron/singulatron/localtron/services/download/types"
)

/*
Starts or resumes a download.
Can resume downloads not found in the JSON statefile.
*/
func (dm *DownloadService) Do(url, downloadDir string) error {
	if downloadDir == "" {
		downloadDir = dm.DefaultFolder
	}

	safeFileName := encodeURLtoFileName(url)
	safeFullFilePath := filepath.Join(downloadDir, safeFileName)
	safePartialFilePath := filepath.Join(downloadDir, safeFileName+".part")

	fullSize, fullFileExists, err := checkFileExistsAndSize(safeFullFilePath)
	if err != nil {
		return err
	}
	partialSize, partialFileExists, err := checkFileExistsAndSize(safePartialFilePath)
	if err != nil {
		return err
	}

	var (
		download *types.Download
		exists   bool
	)

	f := func() error {
		dm.lock.Lock()
		defer dm.lock.Unlock()

		download, exists = dm.downloads[url]

		if !exists {
			if fullFileExists {
				download = &types.Download{
					URL:            url,
					FilePath:       safeFullFilePath,
					Status:         types.DownloadStatusCompleted,
					TotalSize:      fullSize,
					DownloadedSize: fullSize,
				}
				dm.downloads[url] = download
			} else if partialFileExists {

				download = &types.Download{
					URL:            url,
					FilePath:       safeFullFilePath,
					Status:         types.DownloadStatusInProgress,
					DownloadedSize: partialSize,
				}
				dm.downloads[url] = download
			} else {
				download = &types.Download{
					URL:      url,
					FilePath: safeFullFilePath,
					Status:   types.DownloadStatusInProgress,
				}
				dm.downloads[url] = download
			}
		} else {
			// This corrects a potential mismatch between the file size value
			// in the downloads.json and the actual file size which happens
			// if the daemon exists after writing to the file but before reflecting that
			// change in the downloads.json.
			// Search for @transaction-problem in this file
			if partialFileExists {
				download.DownloadedSize = partialSize
			}
			if download.Status == types.DownloadStatusPaused {
				download.Status = types.DownloadStatusInProgress
			}
		}

		return nil
	}
	err = f()
	if err != nil {
		return nil
	}

	dm.markChanged()

	go func() {
		err := dm.downloadFile(download)
		if err != nil {
			logger.Error("Error downlading file",
				slog.String("url", download.URL),
				slog.String("error", err.Error()),
			)
		}
	}()

	return nil
}

func (dm *DownloadService) downloadFile(d *types.Download) error {
	if d.Status == types.DownloadStatusCompleted {
		return nil
	}
	if d.Status == types.DownloadStatusPaused {
		// this should never happen as Do sets this to inProgress
		return fmt.Errorf("cannot download file with status paused")
	}
	out, err := os.OpenFile(d.FilePath+".part", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return errors.Wrap(err, "opening file for download")
	}
	defer out.Close()

	req, err := http.NewRequest("GET", d.URL, nil)
	if err != nil {
		return err
	}

	if d.DownloadedSize > 0 {
		req.Header.Set("Range", fmt.Sprintf("bytes=%d-", d.DownloadedSize))
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	totalSize, _ := getTotalSizeFromHeaders(resp)
	if err != nil {
		fmt.Printf("Error retrieving total size: %v\n", err)
	}

	if resp.StatusCode == http.StatusPartialContent || resp.StatusCode == http.StatusOK {
		buffer := make([]byte, 1024*256) // 256KB buffer
		for {
			if d.Status == types.DownloadStatusPaused {
				return nil
			}
			n, err := resp.Body.Read(buffer)
			if n > 0 {
				_, err = out.Write(buffer[:n])
				// @transaction-problem
				if err != nil {
					return err
				}
				d.DownloadedSize += int64(n)
				if d.TotalSize == 0 && totalSize != 0 {
					d.TotalSize = totalSize
				}
				dm.markChanged()
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				return err
			}
		}
		out.Close()

		err = os.Rename(d.FilePath+".part", d.FilePath)
		if err != nil {
			return err
		}

		d.Status = types.DownloadStatusCompleted
		dm.markChanged()
	} else {
		fmt.Printf("Failed to download: %s, status code: %d\n", d.URL, resp.StatusCode)
	}

	return nil
}

func getTotalSizeFromHeaders(resp *http.Response) (int64, error) {
	// initial downloads without range request headers
	contentLength := resp.Header.Get("Content-Length")
	contentRange := resp.Header.Get("Content-Range")
	if contentLength == "" && contentRange == "" {
		return 0, fmt.Errorf("Content-Length and Content-Range header is missing")
	}

	if contentLength != "" {
		return strconv.ParseInt(contentLength, 10, 64)
	}

	// Content-Range format is typically "bytes start-end/total"
	parts := strings.Split(contentRange, "/")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid Content-Range format")
	}

	totalSize, err := strconv.ParseInt(parts[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing total size from Content-Range: %v", err)
	}

	return totalSize, nil
}
