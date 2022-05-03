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

	fmt.Printf("Servidor rodando na porta %s", config.PORT)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), r))

}
