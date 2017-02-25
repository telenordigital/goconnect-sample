package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	goconnect "github.com/telenordigital/go-connectid"
)

// UserInfoHandlerFunc is a sample implementation of an endpoint
// accessing the session information form the request context.
func UserInfoHandlerFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")

	session := r.Context().Value(goconnect.SessionContext)

	buf, err := json.MarshalIndent(session, "", "  ")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("Got error: %v", err)))
		return
	}
	w.Write(buf)
}
