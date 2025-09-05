package routes

import (
	"net/http"

	auth "github.com/csvitor-dev/go.dev/src/apps/api/controllers"
	"github.com/csvitor-dev/go.dev/src/middlewares"
	"github.com/csvitor-dev/go.dev/types"
)

var authentication = []types.Route{
	{
		Uri:                "/auth/register",
		Method:             http.MethodPost,
		Handler:            auth.Register,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/login",
		Method:             http.MethodPost,
		Handler:            auth.Login,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/password/refresh",
		Method:             http.MethodPost,
		Handler:            auth.RefreshPassword,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:                "/auth/password/recover",
		Method:             http.MethodPost,
		Handler:            auth.RecoverPassword,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/verify-token",
		Method:             http.MethodGet,
		Handler:            auth.VerifyToken,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
	{
		Uri:                "/auth/password/reset",
		Method:             http.MethodPost,
		Handler:            auth.ResetPassword,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger),
	},
}
