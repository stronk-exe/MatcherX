package utils

import (
	// "encoding/json"
	// "fmt"
	// "net/http"

	"github.com/go-playground/validator"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var Validate = validator.New()


// func WriteJSON(w http.ResponseWriter, status int, v any) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(status)
// 	json.NewEncoder(w).Encode(v)
// }

// func WriteError(w http.ResponseWriter, status int, err error) {
// 	WriteJSON(w, status, map[string]string{"error": err.Error()})
// }

// func ParseJSON(r *http.Request, v any) error {
// 	if r.Body == nil {
// 		return fmt.Errorf("missing request body")
// 	}

// 	return json.NewDecoder(r.Body).Decode(v)
// }
