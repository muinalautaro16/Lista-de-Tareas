package config

import (
	"ASTRIC/BackEnd/ASTRIC/helper/env"
	"ASTRIC/BackEnd/shared/usuarios"
	"fmt"
	"net/http"
	"time"
)

// LoadConfigENV configuracion general
/*
---------------------CONFIGURACION GENERAL-----------------------------
*/
func LoadConfigENV() *env.Default {

	var Config env.Default

	Config.APIEnanle = true        //Activa el servidor API
	Config.WSEnable = false        //Activa el servidor WebSocket
	Config.DOCEnable = false       //Activa el servidor de documentacion
	Config.MYSQLHost = "127.0.0.1" //Host de mysql
	Config.MYSQLPass = ""          //Password mysql
	Config.MYSQLPort = 3306        //Puerto de mysql
	Config.MYSQLUser = "root"      // Usuario de mysql
	Config.MONGOURL = ""           //Cadena de conexion de mysql

	return &Config

}

// LoadConfigUser configuracion de usuario
/*
--------------------SECCION CONFIGURACION DE USUARIOS -------------------
*/
func LoadConfigUser() *Usuarios {

	var User Usuarios

	env.ConfigGenral.UserColumnName = "usuario" // Set columna usuario login
	env.ConfigGenral.UserColumnClave = "clave"  //Set culumna de calve de usuario

	//Inicializacion de usuario
	User.Nombre = "Usuario"
	User.Apellido = "Developers"
	User.Mail = "usuario@developers.com"
	User.Activo = true
	User.Admin = true
	User.Clave = "123456"
	User.Reset = false
	User.Usuario = "dev"

	return &User

}

/*
Usuarios Modelo general
swagger:model Usuarios
*/
type Usuarios struct {
	//swagger:ignore
	ID                int       `json:"id,omitempty" gorm:"column:id;primaryKey"`
	CreatedAt         time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt         time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	Nombre            string    `json:"nombre,omitempty" gorm:"column:nombre" validate:"required"`
	Apellido          string    `json:"apellido,omitempty" gorm:"column:apellido" validate:"required"`
	Mail              string    `json:"mail,omitempty" gorm:"column:mail" validate:"required"`
	Usuario           string    `json:"usuario,omitempty" gorm:"column:usuario" validate:"required"`
	Clave             string    `json:"clave,omitempty" gorm:"column:clave" validate:"required"`
	usuarios.Template           //ES OBLIGATORIO
}

// TableName Configuracion de tabla
func (Usuarios) TableName() string {
	return "usuarios"
}

// RutaRaiz Devuelve la web de la ruta raiz
/*
------------------CONFIGURACION DE VISTA INICIAL DE API-----------------------
*/
func RutaRaiz(w http.ResponseWriter, r *http.Request) {

	plantilla := []byte("<h1>API-REST-Full ASTRIC o no</h1>")

	_, _ = w.Write(plantilla)

	fmt.Fprintf(w, "<h1>Todo</h1>")

}
