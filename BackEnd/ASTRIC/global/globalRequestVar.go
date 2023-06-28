package global

import (
	ASTRICmodels "ASTRIC/BackEnd/ASTRIC/pModels"
	"context"
	"database/sql"

	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

// Request Estructura global donde se cargan las BD recibidas por los request
var request ASTRICmodels.GlobalRequest

// responseWS Estructura global de response de request
var responceWS ASTRICmodels.ResponseWS

// SetBDResponseWS setea la bd en el response del websocket
func SetBDResponceWS(bd string) {
	responceWS.BD = bd
}

// SetTagResponseWS seta el tag del response websocket
func SetTagResponceWS(tag string) {
	responceWS.Tag = tag
}

// SetUsuarioResponseWS Setea el usuario del response websocket
func SetUsuarioResponceWS(usuario string) {
	responceWS.Usuario = usuario
}

// SetTablaResponseWS setea la tabla del response websoket
func SetTablaResponceWS(tabla string) {
	responceWS.Tabla = tabla
}

// SetAppResponseWS setea la ruta del websoket
func SetAppResponceWS(app string) {
	responceWS.App = app
}

// SetEndPointResponseWS setea la ruta del websoket
func SetEndPointResponceWS(EndPoint string) {
	responceWS.EndPoint = EndPoint
}

// SetMysql Setea la conexion con la bd solicitada
func SetMysql(mysql *sql.DB) {
	request.Mysql = mysql
}

// SetTKN Setea el token renovado
func SetTKN(tkn string) {
	request.Tkn = tkn
}

// SetMongo Setea la bd de mongo con los datos obtenidos del request
func SetMongo(mongo *mongo.Database) {
	request.Mongo = mongo
}

// SetMysqlORM Setea la bd en la conexion de mysql con ORM
func SetMysqlORM(mysqlOrm *gorm.DB, close context.CancelFunc) {
	request.MysqlORM = mysqlOrm
	request.MysqlORMClose = close
}

// GetResponseWS Retorna el responseWS
func GetResponceWS() ASTRICmodels.ResponseWS {
	return responceWS
}

// GetTKN retorna el token actualizado
func GetTKN() string {
	return request.Tkn
}

// GetRequestVar devuelve el usuario logeado
func GetRequestVar() ASTRICmodels.GlobalRequest {
	return request
}

// GetMongoDB devuelve la conexion a la bd de mongo ya seteada
func GetMongoDB() *mongo.Database {
	return request.Mongo
}

// GetMysql Devuelve la conexion a la bd de mysql ya seteada
func GetMysql() *sql.DB {
	return request.Mysql
}

// GetMysqlORM Devuelve la conexion a  mysql ORM ya setead
func GetMysqlORM() (*gorm.DB, context.CancelFunc) {
	return request.MysqlORM, request.MysqlORMClose
}
