package routes

import (
	"ASTRIC/BackEnd/api/proyectos/controllers"
	"ASTRIC/BackEnd/shared/routes"
)

func rutas(ruta routes.TipoRuta) {
	ruta("/CrearProyecto", controllers.CrearProyecto, "POST", "crea un nuevo proyecto")
	ruta("/EliminarProyecto/{id}", controllers.EliminarProyecto, "DELETE", "eliminar un proyecto")
	ruta("/ObtenerProyectos", controllers.ObtenerProyectos, "GET", "obtener proyectos")
	ruta("/ModificarProyecto", controllers.ModificarProyecto, "PUT", "modificar proyecto")
}
