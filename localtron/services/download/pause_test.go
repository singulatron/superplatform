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

	types "github.com/singulatron/singulatron/localtron/services/download/types"
	firehoseservice "github.com/singulatron/singulatron/localtron/services/firehose"
	"github.com/stretchr/testify/assert"
)

func TestDownloadPauseAndResume(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rangeHeader := r.Header.Get("Range")
		switch rangeHeader {
		case "bytes=0-4":
			w.Header().Set("Content-Range", "bytes 0-4/11")
			w.WriteHeader(http.StatusPartialContent)
			io.WriteString(w, "Hello")
		case "bytes=5-":
			w.Header().Set("Content-Range", "bytes 5-10/11")
			w.WriteHeader(http.StatusPartialContent)
			io.WriteString(w, " world")
		}
	}))
	defer server.Close()

	dir := path.Join(os.TempDir(), "download_test")
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}
	defer os.RemoveAll(dir)

	downloadURL := server.URL + "/file"

	fs, _ := firehoseservice.NewFirehoseService()
	dm, _ := NewDownloadService(fs)
	dm.StateFilePath = path.Join(dir, "downloadPauseResume.json")

	assert.NoError(t, dm.Do(downloadURL, dir))
	time.Sleep(10 * time.Millisecond) // Let it download some data

	assert.NoError(t, dm.Pause(downloadURL))

	d, ok := dm.GetDownload(downloadURL)
	if !ok || d.Status != types.DownloadStatusPaused {
		t.Errorf("Download did not pause correctly")
	}

	assert.NoError(t, dm.Do(downloadURL, dir))

	c := 0
	for {
		time.Sleep(5 * time.Millisecond)
		d, ok := dm.GetDownload(downloadURL)
		if ok && d.Status == types.DownloadStatusCompleted {
			break
		}
		c++
		if c > 30 {
			panic("infinite loop")
		}
	}

	expectedFilePath := filepath.Join(dir, encodeURLtoFileName(downloadURL))
	data, err := os.ReadFile(expectedFilePath)
	if err != nil {
		t.Fatalf("Failed to read downloaded file: %s", err)
	}
	if string(data) != "Hello world" {
		t.Errorf("Downloaded file content incorrect, got: %s, want: %s", data, "Hello world")
	}
}
