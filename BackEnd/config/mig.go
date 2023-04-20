package config

import (
	"gorm.io/gorm"
)

// Migration funcion para realizar migrado de la BD nesecita configuracion
func Migration(con *gorm.DB) error {

	//var actor my.[[MODELO]]

	if err := con.AutoMigrate(

	// Modelos a migrar
	//&[[Puntero de la variable modelo]],
	); err != nil {
		return err
	}

	return nil

}
