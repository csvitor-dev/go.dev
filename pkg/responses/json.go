package responses

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/csvitor-dev/social-media/pkg/types"
)

// Json: writes an JSON body like response message
func Json(w http.ResponseWriter, status types.StatusCode, data any) types.StatusCode {
	w.WriteHeader(int(status))
	w.Header().Set("Content-Type", "application/json")

	if status == http.StatusNoContent || data == nil {
		return status
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalln(err)
		return http.StatusInternalServerError
	}
	return status
}
