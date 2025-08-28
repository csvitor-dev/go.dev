package routes

import (
	"net/http"

	pubs "github.com/csvitor-dev/social-media/src/apps/api/controllers"
	"github.com/csvitor-dev/social-media/src/middlewares"
	"github.com/csvitor-dev/social-media/types"
)

var publications = []types.Route{
	{
		Uri:     "/pubs",
		Method:  http.MethodPost,
		Handler: pubs.Publish,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/pubs",
		Method:  http.MethodGet,
		Handler: pubs.GetAllPubs,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/pubs/{pubId}",
		Method:  http.MethodGet,
		Handler: pubs.GetPubById,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/pubs/{pubId}",
		Method:  http.MethodPut,
		Handler: pubs.UpdatePubById,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
	{
		Uri:     "/pubs/{pubId}",
		Method:  http.MethodDelete,
		Handler: pubs.DeletePubById,
		MiddlewarePipeline: types.NewPipeline().
			With(middlewares.Logger, middlewares.Authorize),
	},
}
