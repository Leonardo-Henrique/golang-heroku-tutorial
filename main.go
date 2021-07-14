package main

import (
	"app/src/config"
	"app/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {

	r := router.GenerateRouter()

	config.LoadEnvVars()

	fmt.Printf("Servidor rodando na porta %d", config.PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.PORT), r))

}
