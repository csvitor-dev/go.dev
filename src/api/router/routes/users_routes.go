package routes

import (
	"net/http"

	users "github.com/csvitor-dev/social-media/src/api/controllers"
)

var baseUri = "/users"
var paramId = baseUri + "/{id}"

var getAllUsers = Route{
	Uri:         baseUri,
	Method:      http.MethodGet,
	Handler:     users.GetAllUsers,
	RequireAuth: false,
}
var getUserById = Route{
	Uri:         paramId,
	Method:      http.MethodGet,
	Handler:     users.GetUserById,
	RequireAuth: false,
}
var createUser = Route{
	Uri:         baseUri,
	Method:      http.MethodPost,
	Handler:     users.CreateUser,
	RequireAuth: false,
}
var updateUser = Route{
	Uri:         paramId,
	Method:      http.MethodPut,
	Handler:     users.UpdateUserById,
	RequireAuth: false,
}
var deleteUser = Route{
	Uri:         paramId,
	Method:      http.MethodDelete,
	Handler:     users.DeleteUserById,
	RequireAuth: false,
}

var UserRoutes = []Route{
	getAllUsers, getUserById,
	createUser, updateUser,
	deleteUser,
}
