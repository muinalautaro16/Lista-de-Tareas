package models

/*
Proyecto Modelo general
swagger:model Proyecto
*/
type Proyecto struct {
	ID        int    `json:"id" gorm:"column:id;primarykey"`
	Nombre    string `json:"nombre" gorm:"column:nombre" validate:"required"`
	Prioridad string `json:"prioridad" gorm:"column:prioridad"`
	Fecha     string `json:"fecha" gorm:"column:fecha"`
}

// TableName tabla de proyectos
func (Proyectos) TableName() string {
	return "Proyectos"
}

// Proyectos modelo de proyectos
type Proyectos struct {
	Proyecto
}
