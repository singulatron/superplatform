/**
 * @license
 * Copyright (c) The Authors (see the AUTHORS file)
 *
 * This source code is licensed under the GNU Affero General Public License v3.0 (AGPLv3).
 * You may obtain a copy of the AGPL v3.0 at https://www.gnu.org/licenses/agpl-3.0.html.
 */
package downloadservice_test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/singulatron/singulatron/localtron/di"
	downloadservice "github.com/singulatron/singulatron/localtron/services/download"
	downloadtypes "github.com/singulatron/singulatron/localtron/services/download/types"
	types "github.com/singulatron/singulatron/localtron/services/download/types"
	usertypes "github.com/singulatron/singulatron/localtron/services/user/types"
	"github.com/stretchr/testify/require"
)

func TestDownloadFile(t *testing.T) {
	fileHostServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	defer fileHostServer.Close()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	token, err := usertypes.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	router = router.SetBearerToken(token)

	err = router.Put(context.Background(), "download-service", "/download", downloadtypes.DownloadRequest{
		URL: fileHostServer.URL,
	}, nil)
	require.NoError(t, err)

	for {
		time.Sleep(5 * time.Millisecond)
		// req := downloadtypes.GetDownloadRequest{}
		rsp := downloadtypes.GetDownloadResponse{}
		err = router.Get(context.Background(), "download-service", fmt.Sprintf("/download/%v", url.PathEscape(fileHostServer.URL)), nil, &rsp)
		require.NoError(t, err)

		if rsp.Exists && rsp.Download.Status == string(types.DownloadStatusCompleted) {
			break
		}
	}

	expectedFilePath := filepath.Join(options.HomeDir, ".singulatron", "downloads", downloadservice.EncodeURLtoFileName(fileHostServer.URL))
	data, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(data))
}

func TestDownloadFileWithPartFile(t *testing.T) {
	fileHostServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rangeHeader := r.Header.Get("Range")
		if rangeHeader != "bytes=5-" {
			t.Errorf("Expected 'bytes=5-' got '%s'", rangeHeader)
		}

		w.Header().Set("Content-Range", "bytes 5-10/11")
		w.WriteHeader(http.StatusPartialContent)
		io.WriteString(w, " world")
	}))
	defer fileHostServer.Close()

	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	token, err := usertypes.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	router = router.SetBearerToken(token)

	downloadURL := fileHostServer.URL + "/file"

	partFilePath := filepath.Join(options.HomeDir, ".singulatron", "downloads", downloadservice.EncodeURLtoFileName(downloadURL)+".part")
	if err := os.WriteFile(partFilePath, []byte("Hello"), 0644); err != nil {
		t.Fatalf("Failed to create part file: %s", err)
	}

	req := downloadtypes.DownloadRequest{
		URL: downloadURL,
	}
	err = router.Put(context.Background(), "download-service", "/download", req, nil)
	require.NoError(t, err)

	for {
		time.Sleep(5 * time.Millisecond)
		req := downloadtypes.GetDownloadRequest{
			Url: downloadURL,
		}
		rsp := downloadtypes.GetDownloadResponse{}
		err = router.Post(context.Background(), "download-service", "/get", req, &rsp)
		require.NoError(t, err)
		if rsp.Exists && rsp.Download.Status == string(types.DownloadStatusCompleted) {
			break
		}
	}

	expectedFilePath := filepath.Join(options.HomeDir, ".singulatron", "downloads", downloadservice.EncodeURLtoFileName(downloadURL))
	data, err := os.ReadFile(expectedFilePath)
	require.NoError(t, err)
	require.Equal(t, "Hello world", string(data))
}

func TestDownloadFileWithFullFile(t *testing.T) {
	hs := &di.HandlerSwitcher{}
	server := httptest.NewServer(hs)
	defer server.Close()

	options := &di.Options{
		Test: true,
		Url:  server.URL,
	}
	universe, starterFunc, err := di.BigBang(options)
	require.NoError(t, err)

	hs.UpdateHandler(universe)
	router := options.Router

	err = starterFunc()
	require.NoError(t, err)

	token, err := usertypes.RegisterUser(router, "someuser", "pw123", "Some name")
	require.NoError(t, err)
	router = router.SetBearerToken(token)

	downloadURL := "full-file"
	fullFilePath := filepath.Join(options.HomeDir, ".singulatron", "downloads", downloadservice.EncodeURLtoFileName(downloadURL))
	require.NoError(t, os.WriteFile(fullFilePath, []byte("Hello world"), 0644))

	req := downloadtypes.DownloadRequest{
		URL: downloadURL,
	}
	err = router.Put(context.Background(), "download-service", "/download", req, nil)
	require.NoError(t, err)

	var (
		d *types.DownloadDetails
	)
	for {
		time.Sleep(5 * time.Millisecond)
		req := downloadtypes.GetDownloadRequest{
			Url: downloadURL,
		}
		rsp := downloadtypes.GetDownloadResponse{}
		err = router.Post(context.Background(), "download-service", "/get", req, &rsp)
		require.NoError(t, err)

		if rsp.Exists && rsp.Download.Status == string(types.DownloadStatusCompleted) {
			d = rsp.Download
			break
		}
	}

	require.Equal(t, int64(11), d.DownloadedBytes)
	require.Equal(t, int64(11), *d.FullFileSize)
}
