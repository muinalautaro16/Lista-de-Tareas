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

	ruta("/CrearEtapa", controllers.CrearEtapa, "POST", "crear una nueva etapa")
	ruta("/EliminarEtapa/{id}", controllers.EliminarEtapa, "DELETE", "eliminar una etapa")
	ruta("/ObtenerEtapas/{id}", controllers.ObtenerEtapas, "GET", "obtener etapas")
	ruta("/ModificarEtapa", controllers.ModificarEtapa, "PUT", "modificar etapa")
}
