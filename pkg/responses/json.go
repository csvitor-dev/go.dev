package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON: writes an JSON body like response message
func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err)
	}
}
