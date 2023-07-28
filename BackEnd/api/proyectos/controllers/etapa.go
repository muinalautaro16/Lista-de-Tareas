package controllers

import (
	"ASTRIC/BackEnd/api/proyectos/models"
	"ASTRIC/BackEnd/shared/db"
	"ASTRIC/BackEnd/shared/ep"
	"encoding/json"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gorilla/mux"
)

// CrearEtapa crea una nueva etapa
func CrearEtapa(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /etapas/CrearEtapa etapas CrearEtapa
	// ---
	// summary: crea una etapa
	//   in: body
	//   description: crea una etapa
	//   schema:
	//     "$ref": "#/definitions/Proyecto"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("etapas/CrearEtapa", w, r)
	res := ep.NewResponse("Crear Etapa", w)

	var etapa models.Etapa

	err := json.NewDecoder(r.Body).Decode(&etapa)
	if err != nil {
		res.ErrSend(err.Error())
		return
	}

	error, valid := ep.ValidateStruct(etapa)
	if valid {
		res.Err(error).DatoSend(etapa)
		return
	}

	if strings.TrimSpace(etapa.Nombre) == "" {
		res.ErrSend("No se pudo crear la etapa.")
		return
	}

	if utf8.RuneCountInString(etapa.Nombre) > 150 {
		res.ErrSend("No se pudo crear la etapa.")
		return
	}

	conexion, close := db.MysqlORM()
	defer close()

	result := conexion.Create(&etapa)

	if result.RowsAffected < 1 {
		res.Err("No se pudo crear la etapa.").DatoSend(etapa)
		return
	}

	res.DatoSend(etapa)
}

// EliminarEtapa elimina una etapa
func EliminarEtapa(w http.ResponseWriter, r *http.Request) {
	// swagger:operation DELETE /etapas/EliminarEtapa Etapas EliminarEtapas
	// ---
	// summary: elimina las etapas
	//   in: body
	//   description: elimina las etapas
	//   schema:
	//     "$ref": "#/definitions/Etapas"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("etapas/EliminarEtapa", w, r)
	res := ep.NewResponse("Eliminar Etapa", w)

	params := mux.Vars(r)

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Delete(&models.Etapa{}, params["id"])

	if result.RowsAffected < 1 {
		res.ErrSend("No se pudo eliminar la etapa.")
		return
	}

	res.DatoSend("La etapa se elimino correctamente")
}

// ObtenerEtapas lista las etapas de un proyecto
func ObtenerEtapas(w http.ResponseWriter, r *http.Request) {
	// swagger:operation GET /etapas/ObtenerEtapas etapas ObtenerEtapas
	// ---
	// summary: obtiene las etapas de un proyecto
	//   in: body
	//   description: obtiene las etapas de un proyecto
	//   schema:
	//     "$ref": "#/definitions/Etapas"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("Etapas/ObtenerEtapas", w, r)
	res := ep.NewResponse("Obtener Etapas", w)

	var etapas []models.Etapa
	idProyecto := mux.Vars(r)

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Where("id_proyecto = ?", idProyecto["id"]).Find(&etapas)

	if result.RowsAffected < 1 {
		res.ErrSend("No se encontraron etapas.")
		return
	}

	res.DatoSend(etapas)
}

// ModificarEtapa modifica una etapa
func ModificarEtapa(w http.ResponseWriter, r *http.Request) {
	// swagger:operation PUT /etapas/ModificarEtapa etapas ModificarEtapa
	// ---
	// summary: modifica una etapa
	//   in: body
	//   description: modifica una etapa
	//   schema:
	//     "$ref": "#/definitions/Proyecto"
	// responses:
	//   default:
	//     description: Respuesta por defecto
	//     schema:
	//       "$ref": "#/definitions/Response"
	defer ep.ErrorControlResponse("etapas/ModificarEtapa", w, r)
	res := ep.NewResponse("Modificar Etapa", w)
	var etapa models.Etapa

	err := json.NewDecoder(r.Body).Decode(&etapa)
	if err != nil {
		res.Err(err.Error())
		return
	}

	error, valid := ep.ValidateStruct(etapa)
	if valid {
		res.Err(error).DatoSend(etapa)
		return
	}

	if strings.TrimSpace(etapa.Nombre) == "" {
		res.ErrSend("No se pudo modificar la etapa.")
		return
	}

	if utf8.RuneCountInString(etapa.Nombre) > 150 {
		res.ErrSend("No se pudo modificar la etapa.")
		return
	}

	conexion, cancel := db.MysqlORM()
	defer cancel()

	result := conexion.Updates(&etapa)

	if result.RowsAffected < 1 {
		res.ErrSend("No se pudo actulizar la etapa")
		return
	}

	res.DatoSend(etapa)
}
