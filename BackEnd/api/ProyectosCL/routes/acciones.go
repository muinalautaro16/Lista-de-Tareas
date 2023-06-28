package routes

import (
	"ASTRIC/BackEnd/api/ProyectosCL/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func acciones(ruta routes.TipoAccion) {
	ruta("/crearProyecto", controllers.CrearProyecto, "POST", "crea un nuevo proyecto")
}
