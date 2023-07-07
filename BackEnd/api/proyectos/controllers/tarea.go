package controllers

import (
	"ASTRIC/BackEnd/api/proyectos/models"
	"ASTRIC/BackEnd/shared/db"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CrearTarea(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Tareas/CrearTarea", w, r)
	res := ep.NewResponse("Crear Tarea", w)

	var tarea models.Tarea

	err := json.NewDecoder(r.Body).Decode(&tarea)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	conexion, cancel := db.MysqlORM()

	result := conexion.Create(&tarea)

	if result.RowsAffected < 1 {
		res.Err("No se pudo crear el proyecto").DatoSend(tarea)
		return
	}

	res.DatoSend(tarea)
	defer cancel()
}

func EliminarTarea(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Tareas/EliminarTarea", w, r)
	res := ep.NewResponse("Eliminar Tarea", w)

	params := mux.Vars(r)

	conexion, cancel := db.MysqlORM()

	result := conexion.Delete(&models.Tarea{}, params["id"])
	if result.RowsAffected < 1 {
		res.ErrSend("No se pudo eliminar la tarea")
		return
	}

	res.DatoSend("Tarea eliminada")
	defer cancel()
}

func ObtenerTareas(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Tareas/ObtenerTareas", w, r)
	res := ep.NewResponse("Obtener Tareas", w)

	var tareas []models.Tarea

	conexion, cancel := db.MysqlORM()

	result := conexion.Find(&tareas)

	if result.Error != nil {
		res.Err("No se pudieron obtener las tareas de la etapa")
		return
	}
	defer cancel()

	json.NewEncoder(w).Encode(tareas)
}

func ModificarTarea(w http.ResponseWriter, r *http.Request) {
	defer ep.ErrorControlResponse("Tareas/ModificarTarea", w, r)
	res := ep.NewResponse("Modificar Tareas", w)

	var tarea models.Tarea

	err := json.NewDecoder(r.Body).Decode(&tarea)
	if err != nil {
		res.Err(err.Error())
		return
	}

	conexion, cancel := db.MysqlORM()

	result := conexion.Updates(&tarea)
	if result.RowsAffected < 1 {
		res.ErrSend("No se pudo moficiar la tarea")
		return
	}

	res.DatoSend("Tarea modificada")
	defer cancel()
}
