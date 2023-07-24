package controllers

import (
	"ASTRIC/BackEnd/shared/ep"
	"net/http"
)

// Prueba prueba
func Prueba(w http.ResponseWriter, r *http.Request) {

	ep.ErrorControlResponse("Prueba", w, r)

	res := ep.NewResponse("Prueba", w)

	res.DatoSend("Prueba")

}
