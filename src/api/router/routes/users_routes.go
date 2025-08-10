package routes

import (
	"net/http"

	users "github.com/csvitor-dev/social-media/src/api/controllers"
)

var userRoutes = []Route{
	{
		Uri:         "/users",
		Method:      http.MethodGet,
		Handler:     users.GetAllUsers,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{id}",
		Method:      http.MethodGet,
		Handler:     users.GetUserById,
		RequireAuth: false,
	},
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Handler:     users.CreateUser,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{id}",
		Method:      http.MethodPut,
		Handler:     users.UpdateUserById,
		RequireAuth: false,
	},
	{
		Uri:         "/users/{id}",
		Method:      http.MethodDelete,
		Handler:     users.DeleteUserById,
		RequireAuth: false,
	},
}
