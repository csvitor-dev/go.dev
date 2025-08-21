package routes

import (
	"net/http"

	auth "github.com/csvitor-dev/social-media/src/apps/api/controllers"
	"github.com/csvitor-dev/social-media/src/middlewares"
	"github.com/csvitor-dev/social-media/types"
)

var authentication = []types.Route{
	{
		Uri:     "/auth/register",
		Method:  http.MethodPost,
		Handler: auth.Register,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger),
	},
	{
		Uri:     "/auth/login",
		Method:  http.MethodPost,
		Handler: auth.Login,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger),
	},
	{
		Uri:     "/auth/password/refresh",
		Method:  http.MethodPost,
		Handler: auth.RefreshPassword,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:                "/auth/password/recover",
		Method:             http.MethodPost,
		Handler:            auth.RecoverPassword,
		MiddlewarePipeline: types.NewPipeline(),
	},
	{
		Uri:                "/auth/password/reset",
		Method:             http.MethodGet,
		Handler:            auth.ValidateResetPasswordToken,
		MiddlewarePipeline: types.NewPipeline(),
	},
	{
		Uri:                "/auth/password/reset",
		Method:             http.MethodPost,
		Handler:            auth.ResetPassword,
		MiddlewarePipeline: types.NewPipeline(),
	},
}
