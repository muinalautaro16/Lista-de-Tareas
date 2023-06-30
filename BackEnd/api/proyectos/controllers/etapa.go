package controllers

import (
	"ASTRIC/BackEnd/api/proyectos/models"
	"ASTRIC/BackEnd/shared/db"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CrearEtapa(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Etapas/CrearEtapa", w, r)
	res := ep.NewResponse("Crear Etapa", w)

	var etapa models.Etapa

	err := json.NewDecoder(r.Body).Decode(&etapa)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	conexion, cancel := db.MysqlORM()

	result := conexion.Create(&etapa)

	if result.RowsAffected < 1 {
		res.Err("No se pudo crear la etapa.").DatoSend(etapa)
		return
	}

	res.DatoSend(etapa)
	defer cancel()
}

func EliminarEtapa(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Etapas/Eliminar", w, r)
	res := ep.NewResponse("Eliminar Etapa", w)

	params := mux.Vars(r)

	conexion, cancel := db.MysqlORM()

	result := conexion.Delete(&models.Etapa{}, params["id"])

	if result.Error != nil {
		res.Err("No se pudo eliminar la etapa.")
		return
	}

	defer cancel()
}

func ObtenerEtapas(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Etapas/ObtenerEtapas", w, r)
	res := ep.NewResponse("Obtener Etapas", w)

	var etapas []models.Etapa

	conexion, cancel := db.MysqlORM()

	result := conexion.Find(&etapas)

	if result.Error != nil {
		res.Err("No se puedo obtener las etapas del proyecto")
		return
	}

	defer cancel()

	json.NewEncoder(w).Encode(etapas)
}

func ModificarEtapa(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Etapas/ObtenerEtapas", w, r)
	res := ep.NewResponse("Obtener Etapas", w)
	var etapa models.Etapa

	err := json.NewDecoder(r.Body).Decode(&etapa)
	if err != nil {
		res.Err(err.Error())
		return
	}

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Updates(&etapa)

	if result.RowsAffected < 1 {
		res.Err("No se puede actualizar la etapa").DatoSend(etapa)
		return
	}

	res.DatoSend(etapa)
}
