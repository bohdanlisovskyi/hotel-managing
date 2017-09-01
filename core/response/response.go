package response

import (
	"encoding/json"
	"net/http"
)

// encode incoming interface to json
func New(w http.ResponseWriter, message interface{}) {
	json.NewEncoder(w).Encode(message)
}
