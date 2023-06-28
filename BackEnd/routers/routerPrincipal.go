package routers

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	"ASTRIC/BackEnd/ASTRIC/middleware"
	rProyecto "ASTRIC/BackEnd/api/proyectos/routes"

	"github.com/gorilla/mux"
)

// RutasPrincipales Manejador de rutas principales, donde se declaran los prefijos
func RutasPrincipales(rout *mux.Router) {

	rutaProyecto := routes.NewPrefix(rout, "/Proyectos", "Modelo de Proyectos")

	rutaProyecto.Use(middleware.ProcesarRutas)
	routerProyecto := routes.NewRouter(rutaProyecto)
	rProyecto.RutasModul(routerProyecto)

}
