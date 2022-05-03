package routes

import (
	"app/src/controllers"
	"net/http"
)

var UserRoute = []DefaultRoute{
	{
		URI:    "/",
		Method: http.MethodGet,
		Func:   controllers.Welcome,
	},
}
