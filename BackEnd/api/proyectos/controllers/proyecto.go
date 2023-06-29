package controllers

import (
	"ASTRIC/BackEnd/api/proyectos/models"
	"ASTRIC/BackEnd/shared/db"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"

	"net/http"

	"github.com/gorilla/mux"
)

func CrearProyecto(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Proyecto/CrearProyecto", w, r)
	res := ep.NewResponse("Crear Proyecto", w)

	var proyecto models.Proyecto

	err := json.NewDecoder(r.Body).Decode(&proyecto)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	// errores, invalido := ep.ValidateStruct(proyecto)
	// if invalido {
	// 	res.Err(errores).DatoSend(proyecto)
	// 	return
	// }

	conexion, cancel := db.MysqlORM()

	result := conexion.Create(&proyecto)

	if result.RowsAffected < 1 {
		res.Err("No se puedo crear el proyecto.").DatoSend(proyecto)
		return
	}

	res.DatoSend(proyecto)
	defer cancel()
}

func EliminarProyecto(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Proyecto/EliminarProyecto", w, r)
	res := ep.NewResponse("Eliminar Proyecto", w)

	params := mux.Vars(r)

	conexion, cancel := db.MysqlORM()

	result := conexion.Delete(&models.Proyecto{}, params["id"])

	if result.Error != nil {
		res.Err("No se puede eliminar el Proyecto")
		return
	}

	defer cancel()
}

func ObtenerProyectos(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Proyecto/ObtenerProyectos", w, r)
	res := ep.NewResponse("Obtener Proyectos", w)

	var proyectos []models.Proyecto

	conexion, cancel := db.MysqlORM()

	result := conexion.Find(&proyectos)

	if result.Error != nil {
		res.Err("No se pudiern obtener proyectos")
		return
	}

	defer cancel()

	json.NewEncoder(w).Encode(proyectos)
}

func ModificarProyecto(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Proyecto/ObtenerProyectos", w, r)
	res := ep.NewResponse("Obtener Proyectos", w)
	var proyecto models.Proyecto

	err := json.NewDecoder(r.Body).Decode(&proyecto)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Updates(&proyecto)

	if result.RowsAffected < 1 {
		res.Err("No se pudo actualizar el proyecto").DatoSend(proyecto)
	}

	res.DatoSend(proyecto)
}
