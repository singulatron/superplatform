/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadservice_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/di"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	types "github.com/singulatron/singulatron/localtron/services/download/types"
	"github.com/stretchr/testify/require"
)

func TestDownloadFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rangeHeader := r.Header.Get("Range")
		if rangeHeader != "" {
			w.Header().Set("Content-Range", "bytes 0-10/11")
			w.WriteHeader(http.StatusPartialContent)
			io.WriteString(w, "Hello world")
		} else {
			w.WriteHeader(http.StatusOK)
			io.WriteString(w, "Hello world")
		}
	}))
	defer server.Close()

	dir := path.Join(os.TempDir(), "download_test")
	require.NoError(t, os.MkdirAll(dir, 0755))

	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	dm := universe.DownloadService

	dm.StateFilePath = path.Join(dir, "downloadFile.json")
	require.NoError(t, dm.Do(server.URL, dir))

	for {
		time.Sleep(5 * time.Millisecond)
		d, ok := dm.GetDownload(server.URL)
		if ok && d.Status == types.DownloadStatusCompleted {
			break
		}
	}

	expectedFilePath := filepath.Join(dir, downloadservice.EncodeURLtoFileName(server.URL))
	data, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(data))
}

func TestDownloadFileWithPartFile(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rangeHeader := r.Header.Get("Range")
		if rangeHeader != "bytes=5-" {
			t.Errorf("Expected 'bytes=5-' got '%s'", rangeHeader)
		}

		w.Header().Set("Content-Range", "bytes 5-10/11")
		w.WriteHeader(http.StatusPartialContent)
		io.WriteString(w, " world")
	}))
	defer server.Close()

	dir := path.Join(os.TempDir(), "download_test")
	require.NoError(t, os.MkdirAll(dir, 0755))

	downloadURL := server.URL + "/file"
	partFilePath := filepath.Join(dir, downloadservice.EncodeURLtoFileName(downloadURL)+".part")
	if err := os.WriteFile(partFilePath, []byte("Hello"), 0644); err != nil {
		t.Fatalf("Failed to create part file: %s", err)
	}

	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	dm := universe.DownloadService

	dm.StateFilePath = path.Join(dir, "downloadFilePartial.json")

	require.NoError(t, dm.Do(downloadURL, dir))

	for {
		time.Sleep(5 * time.Millisecond)
		d, ok := dm.GetDownload(downloadURL)
		if ok && d.Status == types.DownloadStatusCompleted {
			break
		}
	}

	expectedFilePath := filepath.Join(dir, downloadservice.EncodeURLtoFileName(downloadURL))
	data, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(data))
}

func TestDownloadFileWithFullFile(t *testing.T) {
	dir := path.Join(os.TempDir(), "download_test")
	require.NoError(t, os.MkdirAll(dir, 0755))

	downloadURL := "full-file"
	fullFilePath := filepath.Join(dir, downloadservice.EncodeURLtoFileName(downloadURL))
	require.NoError(t, os.WriteFile(fullFilePath, []byte("Hello world"), 0644))

	universe, err := di.BigBang(di.UniverseOptions{
		Test: true,
	})
	require.NoError(t, err)
	dm := universe.DownloadService

	dm.StateFilePath = path.Join(dir, "downloadFileFull.json")
	require.NoError(t, dm.Do(downloadURL, dir))

	var (
		d  *types.Download
		ok bool
	)
	for {
		time.Sleep(5 * time.Millisecond)
		d, ok = dm.GetDownload(downloadURL)
		if ok && d.Status == types.DownloadStatusCompleted {
			break
		}
	}

	require.Equal(t, int64(11), d.DownloadedSize)
	require.Equal(t, int64(11), d.TotalSize)
}
