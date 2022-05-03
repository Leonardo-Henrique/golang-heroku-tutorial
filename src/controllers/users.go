package controllers

import "net/http"

func Welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá! Seja muito bem vindo(a)!"))
}
