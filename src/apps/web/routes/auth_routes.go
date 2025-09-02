package routes

import (
	"net/http"

	"github.com/csvitor-dev/social-media/src/apps/web/controllers"
	"github.com/csvitor-dev/social-media/src/middlewares"
	"github.com/csvitor-dev/social-media/types"
	utils "github.com/csvitor-dev/social-media/utils/http"
)

var auth = []types.Route{
	{
		Uri:    "/",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			utils.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		},
		MiddlewarePipeline: types.NewPipeline().With(middlewares.Logger),
	},
	{
		Uri:                "/auth/login",
		Method:             http.MethodGet,
		Handler:            controllers.GetLoginView,
		MiddlewarePipeline: types.NewPipeline().With(middlewares.Logger),
	},
	{
		Uri:                "/auth/reset-password",
		Method:             http.MethodGet,
		Handler:            controllers.FetchApiForTokenValidation,
		MiddlewarePipeline: types.NewPipeline().With(middlewares.Logger),
	},
}
