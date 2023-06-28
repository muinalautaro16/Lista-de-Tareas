package routers

import (
	"ASTRIC/BackEnd/ASTRIC/helper/routes"
	rProyecto "ASTRIC/BackEnd/api/ProyectosCL/routes"

	"github.com/gorilla/mux"
)

// RutasPrincipales Manejador de rutas principales, donde se declaran los prefijos
func RutasPrincipales(rout *mux.Router) {
	//Rutas de juzgados
	routeProyecto := routes.NewPrefix(rout, "/proyecto", "Modulo de proyectos")
	//ruotersPrueba.Use(database.ChequeBD)
	// ruottotems.Use(middleware.ProcesarRutas)

	routerProyecto := routes.NewRouter(routeProyecto)
	rProyecto.RutasModul(routerProyecto)
}
