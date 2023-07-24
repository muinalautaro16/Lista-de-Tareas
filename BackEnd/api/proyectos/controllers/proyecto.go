package controllers

import (
	"ASTRIC/BackEnd/api/proyectos/models"
	"ASTRIC/BackEnd/shared/db"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// CrearProyecto crea un proyecto
func CrearProyecto(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /proyectos/CrearProyecto Proyectos CrearProyecto
	// ---
	// summary: crea un proyecto
	//   in: body
	//   description: crea un proyecto
	//   schema:
	//     "$ref": "#/definitions/Proyecto"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("proyecto/CrearProyecto", w, r)
	res := ep.NewResponse("Crear Proyecto", w)

	var proyecto models.Proyecto

	err := json.NewDecoder(r.Body).Decode(&proyecto)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	error, valid := ep.ValidateStruct(proyecto)
	if valid {
		res.Err(error).DatoSend(proyecto)
		return
	}

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Create(&proyecto)

	if result.RowsAffected < 1 {
		res.Err("No se puedo crear el proyecto.").DatoSend(proyecto)
		return
	}

	res.DatoSend(proyecto)
}

// EliminarProyecto elimina un proyecto
func EliminarProyecto(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /proyectos/EliminarProyecto Proyectos EliminarProyecto
	// ---
	// summary: elimina un proyecto
	//   in: body
	//   description: elimina un proyecto
	//   schema:
	//     "$ref": "#/definitions/Proyecto"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("proyecto/EliminarProyecto", w, r)
	res := ep.NewResponse("Eliminar Proyecto", w)

	params := mux.Vars(r)

	conexion, cancel := db.MysqlORM()

	result := conexion.Delete(&models.Proyecto{}, params["id"])
	if result.RowsAffected < 1 {
		res.ErrSend("No se pudo eliminar el Proyecto")
		return
	}

	resultado := conexion.Where("id_proyecto = ?", params["id"]).Delete(&models.Etapa{})
	if resultado.Error != nil {
		res.ErrSend(resultado.Error)
		return
	}

	res.DatoSend("Proyecto eliminado")
	defer cancel()
}

// ObtenerProyectos lista los proyectos
func ObtenerProyectos(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /proyectos/ObtenerProyectos Proyectos ObtenerProyectos
	// ---
	// summary: lista los proyectos
	//   in: body
	//   description: lista los proyectos
	//   schema:
	//     "$ref": "#/definitions/Proyecto"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("proyecto/ObtenerProyectos", w, r)
	res := ep.NewResponse("Obtener Proyectos", w)

	var proyectos []models.Proyecto

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Find(&proyectos)

	if result.Error != nil {
		res.Err("No se pudieron obtener proyectos")
		return
	}

	res.DatoSend(proyectos)
}

// ModificarProyecto modifica un proyecto
func ModificarProyecto(w http.ResponseWriter, r *http.Request) {
	// swagger:operation UPDATE /proyectos/ModificarProyecto Proyectos ModificarProyecto
	// ---
	// summary: modifica un proyecto
	//   in: body
	//   description: modifica un proyecto
	//   schema:
	//     "$ref": "#/definitions/Proyecto"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("proyecto/ModificarProyecto", w, r)
	res := ep.NewResponse("Modificar Proyecto", w)
	var proyecto models.Proyecto

	err := json.NewDecoder(r.Body).Decode(&proyecto)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	error, valid := ep.ValidateStruct(proyecto)
	if valid {
		res.Err(error).DatoSend(proyecto)
		return
	}

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Updates(&proyecto)

	if result.RowsAffected < 1 {
		res.ErrSend("No se pudo actualizar el proyecto")
		return
	}

	res.DatoSend(proyecto)
}
