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
package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

// GzipDecodeMiddleware checks if the request has Content-Encoding as gzip
// and if so, replaces the request body with the gzip-decompressed version.
func GzipDecodeMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Content-Encoding"), "gzip") {
			gzReader, err := gzip.NewReader(r.Body)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			defer gzReader.Close()

			// Replace request body with decompressed body
			r.Body = io.NopCloser(gzReader)
		} else if strings.Contains(r.Header.Get("Content-Encoding"), "br") {
			brReader := brotli.NewReader(r.Body)

			// Replace request body with decompressed body
			r.Body = io.NopCloser(brReader)
		}

		next(w, r)
	}
}
