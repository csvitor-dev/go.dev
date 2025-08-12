package routes

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/api/controllers"
	"github.com/csvitor-dev/social-media/src/api/middlewares"
)

var authRoutes = []Route{
	{
		Uri:            "/auth/login",
		Method:         http.MethodPost,
		Handler:        controllers.Login,
		MiddlewareTags: middlewares.NewTags(),
	},
}
