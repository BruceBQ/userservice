package web

import "net/http"

type responseWriterWrapper struct {
	http.ResponseWriter
	statusCode        int
	statusCodeWritten bool
}

func newWrappedWriter(original http.ResponseWriter) *responseWriterWrapper {
	return &responseWriterWrapper{
		ResponseWriter:    original,
		statusCodeWritten: false,
	}
}

func (rw *responseWriterWrapper) StatusCode() int {
	return rw.statusCode
}

func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
	rw.statusCode = statusCode
	rw.statusCodeWritten = true
	rw.ResponseWriter.WriteHeader(statusCode)
}

func (rw *responseWriterWrapper) Write(data []byte) (int, error) {
	if !rw.statusCodeWritten {
		rw.statusCode = http.StatusOK
	}
	return rw.ResponseWriter.Write(data)
}
