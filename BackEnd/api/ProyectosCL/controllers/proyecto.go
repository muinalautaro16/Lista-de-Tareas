package controllers

import (
	"ASTRIC/BackEnd/api/ProyectosCL/models"
	"ASTRIC/BackEnd/shared/db"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"net/http"
)

func CrearProyecto(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("proyecto/crearProyecto", w, r)
	res := ep.NewResponse("Crear Proyecto", w)

	var proyecto models.Proyecto

	err := json.NewDecoder(r.Body).Decode(&proyecto)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	errores, invalido := ep.ValidateStruct(proyecto)
	if invalido {
		res.Err(errores).DatoSend(proyecto)
		return
	}

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Create(proyecto)

	if result.RowsAffected < 1 {
		res.Err("No se puedo crear el proyecto.").DatoSend(proyecto)
		return
	}

	res.DatoSend(proyecto)
}
