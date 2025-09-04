package routes

import (
	"net/http"

	users "github.com/csvitor-dev/go.dev/src/apps/api/controllers"
	"github.com/csvitor-dev/go.dev/src/middlewares"
	"github.com/csvitor-dev/go.dev/types"
)

var user = []types.Route{
	{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: users.GetAllUsers,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/users/me",
		Method:  http.MethodGet,
		Handler: users.GetAuthUser,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodGet,
		Handler: users.GetUserById,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: users.UpdateUserById,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: users.DeleteUserById,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
}
