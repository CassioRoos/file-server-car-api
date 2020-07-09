package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipMiddleware struct{}

// The standard way to create a middleware function
func (g *GzipMiddleware) GzipMiddleWare(next http.Handler) http.Handler {
	// cast the function to HandlerFunc
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		// try to get gzip from the header
		// If accept encoding not contains gzip, will do as always does
		if !strings.ContainsAny(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(rw, r)
			return
		}
		// If gzip is found, will create a gzip struct and wrap the ResponseWriter
		gwr := NewWrappedResponse(rw)
		gwr.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(gwr, r)
		defer gwr.Flush()
	})
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

// Since gzip uses the same signature, we can wrap ResponseWrite with gzip
// when the header comes, it will GZIP the content and return it
func NewWrappedResponse(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)
	return &WrappedResponseWriter{
		rw: rw,
		gw: gw,
	}
}

// Return the same header. This is a nice way to wrap functionalities.
func (w *WrappedResponseWriter) Header() http.Header {
	return w.rw.Header()
}

// Write the same header. This is a nice way to wrap functionalities.
func (w *WrappedResponseWriter) WriteHeader(statuscode int) {
	w.rw.WriteHeader(statuscode)
}

// Now we can use the gzip to write the content, the resp of the process remains the same
func (w *WrappedResponseWriter) Write(d []byte) (int, error) {
	return w.gw.Write(d)
}

// Write everything that was waiting to be written
func (w *WrappedResponseWriter) Flush() {
	w.gw.Flush()
	w.gw.Close()
}
