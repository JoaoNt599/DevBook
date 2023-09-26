package routes

import "net/http"

var userRoutes = []Route{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuthentication: false,
	},

	{
		URI:    "/users",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuthentication: false,
	},

	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuthentication: false,
	},

	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuthentication: false,
	},

	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Function: func(w http.ResponseWriter, r *http.Request) {

		},
		RequiresAuthentication: false,
	},
}
