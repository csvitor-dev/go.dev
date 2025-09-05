package routes

import (
	"net/http"

	followers "github.com/csvitor-dev/go.dev/src/apps/api/controllers"
	"github.com/csvitor-dev/go.dev/src/middlewares"
	"github.com/csvitor-dev/go.dev/types"
)

var follower = []types.Route{
	{
		Uri:                "/follow/{userId}",
		Method:             http.MethodPost,
		Handler:            followers.Follow,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:                "/unfollow/{userId}",
		Method:             http.MethodPost,
		Handler:            followers.Unfollow,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:                "/followers/{userId}",
		Method:             http.MethodGet,
		Handler:            followers.GetFollowers,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:                "/following/{userId}",
		Method:             http.MethodGet,
		Handler:            followers.GetFollowing,
		MiddlewarePipeline: types.NewPipeline(middlewares.Logger, middlewares.Authorize),
	},
}
