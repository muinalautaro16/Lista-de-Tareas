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
	ruta("/ObtenerEtapas", controllers.ObtenerEtapas, "GET", "obtener etapas")
	ruta("/ModificarEtapa", controllers.ModificarEtapa, "PUT", "modificar etapa")

	ruta("/CrearTarea", controllers.CrearTarea, "POST", "crear una nueva tarea")
	ruta("/EliminarTarea/{id}", controllers.EliminarTarea, "DELETE", "eliminar una tarea")
	ruta("/ObtenerTareas", controllers.ObtenerTareas, "GET", "obtener tareas")
	ruta("/ModificarTarea", controllers.ModificarTarea, "PUT", "modificar tarea")
}
