package routes

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/api/controllers"
)

var authRoutes = []Route{
	{
		Uri:         "/auth/login",
		Method:      http.MethodPost,
		Handler:     controllers.Login,
		RequireAuth: false,
	},
}
