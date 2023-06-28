package routes

import (
	"ASTRIC/BackEnd/api/proyectos/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func rutas(ruta routes.TipoRuta) {
	ruta("/CrearProyecto", controllers.CrearProyecto, "POST", "crea un nuevo proyecto")
}
