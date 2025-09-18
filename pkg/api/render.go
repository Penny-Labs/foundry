package api

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// RenderJson marshals 'v' to JSON, automatically escaping HTML, setting the
// Content-Type as application/json, and sending the status code header.
func RenderJson(w http.ResponseWriter, v interface{}, statusCode int) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, _ = w.Write(buf.Bytes())
}

func JsonError(w http.ResponseWriter, message string, code int) {
	w.Header().Del("Content-Length")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	RenderJson(w, map[string]string{"error": message}, code)
}
