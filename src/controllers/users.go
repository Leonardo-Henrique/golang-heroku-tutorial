package controllers

import "net/http"

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando usuarios"))
}

func Welcome(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Olá! Seja muito bem vindo(a)!"))
}