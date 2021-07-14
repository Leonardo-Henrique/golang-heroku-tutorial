package router

import (
	"app/src/router/routes"

	"github.com/gorilla/mux"
)

func GenerateRouter() *mux.Router {
	r := mux.NewRouter()

	return routes.InjectRoutes(r)
}
