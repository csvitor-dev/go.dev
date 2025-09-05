package routes

import (
	"net/http"

	"github.com/csvitor-dev/go.dev/src/apps/web/controllers/actions"
	"github.com/csvitor-dev/go.dev/src/apps/web/controllers/views"
	"github.com/csvitor-dev/go.dev/src/middlewares"
	"github.com/csvitor-dev/go.dev/types"
	utils "github.com/csvitor-dev/go.dev/utils/http"
)

var authViews = []types.Route{
	{
		Uri:    "/",
		Method: http.MethodGet,
		Handler: func(w http.ResponseWriter, r *http.Request) {
			utils.Redirect(w, r, "/auth/login", http.StatusSeeOther)
		},
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/login",
		Method:             http.MethodGet,
		Handler:            views.GetLoginView,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/register",
		Method:             http.MethodGet,
		Handler:            views.GetRegisterView,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/forgot-password",
		Method:             http.MethodGet,
		Handler:            views.GetForgotPasswordView,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/reset-password",
		Method:             http.MethodGet,
		Handler:            views.GetResetPasswordView,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
}

var authActions = []types.Route{
	{
		Uri:                "/auth/register",
		Method:             http.MethodPost,
		Handler:            actions.RegisterUserAction,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
}
