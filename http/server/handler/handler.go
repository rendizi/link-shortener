package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func isValid(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path[len("/"):]

	segments := strings.Split(path, "/")

	if len(segments) > 0 && segments[0] != "" {
		fmt.Println("Requested segment:", segments[0])
	} else {
		fmt.Println("No segment provided after the base URL")
	}
}
