package router

import (
	"app/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	return routes.InjectRoutes(mux.NewRouter())
}
