package routes

import (
	"net/http"

	c "github.com/csvitor-dev/social-media/api/src/controllers"
)

var baseURI = "/users"
var paramID = baseURI + "/{id}"

var getAllUsers = Route{
	URI: baseURI,
	Method: http.MethodGet,
	Handler: c.GetAllUsers,
	RequireAuth: false,
}
var getUserByID = Route{
	URI: paramID,
	Method: http.MethodGet,
	Handler: c.GetUserByID,
	RequireAuth: false,
}
var createUser = Route{
	URI: baseURI,
	Method: http.MethodPost,
	Handler: c.CreateUser,
	RequireAuth: false,
}
var updateUser = Route{
	URI: paramID,
	Method: http.MethodPut,
	Handler: c.UpdateUserByID,
	RequireAuth: false,
}
var deleteUser = Route{
	URI: paramID,
	Method: http.MethodDelete,
	Handler: c.DeleteUserByID,
	RequireAuth: false,
}

var UserRoutes = []Route{
	getAllUsers, getUserByID,
	createUser, updateUser,
	deleteUser,
}