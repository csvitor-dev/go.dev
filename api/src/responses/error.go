package responses

import "net/http"

func Error(w http.ResponseWriter, status int, errs []error) {
	var hook []string

	for _, v := range errs {
		hook = append(hook, v.Error())
	}
	JSON(w, status, struct {
		Errors []string `json:"errors"`
	}{
		Errors: hook,
	})
}