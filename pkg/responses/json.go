package responses

import (
	"encoding/json"
	"log"
	"net/http"

	utils "github.com/csvitor-dev/social-media/utils/http"
)

// Json: writes an JSON body like response message
func Json(w http.ResponseWriter, status int, data any) {
	utils.WriteStatus(w, status)
	w.Header().Set("Content-Type", "application/json")

	if status == http.StatusNoContent || data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err)
	}
}
