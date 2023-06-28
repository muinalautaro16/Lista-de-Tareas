package routes

import "ASTRIC/BackEnd/ASTRIC/helper/routes"

func RutasModul(rutastruct routes.RoutersStruct) {
	rutas(rutastruct.Ruta)
	acciones(rutastruct.Accion)
}
