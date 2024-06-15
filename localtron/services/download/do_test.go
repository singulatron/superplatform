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
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/datastore/localstore"
	configservice "github.com/singulatron/singulatron/localtron/services/config"
	types "github.com/singulatron/singulatron/localtron/services/download/types"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	userservice "github.com/singulatron/singulatron/localtron/services/user"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/stretchr/testify/assert"
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
	assert.NoError(t, os.MkdirAll(dir, 0755))

	cs, _ := configservice.NewConfigService()
	us, _ := userservice.NewUserService(cs, localstore.NewLocalStore[*usertypes.User](""),
		localstore.NewLocalStore[*usertypes.Role](""),
		localstore.NewLocalStore[*usertypes.AuthToken](""),
		localstore.NewLocalStore[*usertypes.Permission](""))
	fs, _ := firehoseservice.NewFirehoseService(us)
	dm, _ := NewDownloadService(fs, us)
	dm.StateFilePath = path.Join(dir, "downloadFile.json")
	assert.NoError(t, dm.Do(server.URL, dir))

	for {
		time.Sleep(5 * time.Millisecond)
		d, ok := dm.GetDownload(server.URL)
		if ok && d.Status == types.DownloadStatusCompleted {
			break
		}
	}

	expectedFilePath := filepath.Join(dir, encodeURLtoFileName(server.URL))
	data, err := os.ReadFile(expectedFilePath)
	assert.NoError(t, err)
	assert.Equal(t, "Hello world", string(data))
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
	assert.NoError(t, os.MkdirAll(dir, 0755))

	downloadURL := server.URL + "/file"
	partFilePath := filepath.Join(dir, encodeURLtoFileName(downloadURL)+".part")
	if err := os.WriteFile(partFilePath, []byte("Hello"), 0644); err != nil {
		t.Fatalf("Failed to create part file: %s", err)
	}

	cs, _ := configservice.NewConfigService()
	us, _ := userservice.NewUserService(cs, localstore.NewLocalStore[*usertypes.User](""),
		localstore.NewLocalStore[*usertypes.Role](""),
		localstore.NewLocalStore[*usertypes.AuthToken](""),
		localstore.NewLocalStore[*usertypes.Permission](""))
	fs, _ := firehoseservice.NewFirehoseService(us)
	dm, _ := NewDownloadService(fs, us)
	dm.StateFilePath = path.Join(dir, "downloadFilePartial.json")

	assert.NoError(t, dm.Do(downloadURL, dir))

	for {
		time.Sleep(5 * time.Millisecond)
		d, ok := dm.GetDownload(downloadURL)
		if ok && d.Status == types.DownloadStatusCompleted {
			break
		}
	}

	expectedFilePath := filepath.Join(dir, encodeURLtoFileName(downloadURL))
	data, err := os.ReadFile(expectedFilePath)
	assert.NoError(t, err)
	assert.Equal(t, "Hello world", string(data))
}

func TestDownloadFileWithFullFile(t *testing.T) {
	dir := path.Join(os.TempDir(), "download_test")
	assert.NoError(t, os.MkdirAll(dir, 0755))

	downloadURL := "full-file"
	fullFilePath := filepath.Join(dir, encodeURLtoFileName(downloadURL))
	assert.NoError(t, os.WriteFile(fullFilePath, []byte("Hello world"), 0644))

	cs, _ := configservice.NewConfigService()
	us, _ := userservice.NewUserService(cs,
		localstore.NewLocalStore[*usertypes.User](""),
		localstore.NewLocalStore[*usertypes.Role](""),
		localstore.NewLocalStore[*usertypes.AuthToken](""),
		localstore.NewLocalStore[*usertypes.Permission](""),
	)
	fs, _ := firehoseservice.NewFirehoseService(us)
	dm, _ := NewDownloadService(fs, us)
	dm.StateFilePath = path.Join(dir, "downloadFileFull.json")
	assert.NoError(t, dm.Do(downloadURL, dir))

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

	assert.Equal(t, int64(11), d.DownloadedSize)
	assert.Equal(t, int64(11), d.TotalSize)
}
