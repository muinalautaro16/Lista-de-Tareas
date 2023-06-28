package models

type Proyecto struct {
	ID           int64  `json:"id" gorm:"column:ID"`
	Nombre       string `json:"nombre" gorm:"column:NOMBRE" validate:"required"`
	Prioridad    string `json:"prioridad" gorm:"column:prioridad" validate:"required"`
	Fecha_inicio string `json:"fecha_inicio" gorm:"column:fecha_inicio"`
	Fecha_fin    string `json:"fecha_fin" gorm:"column:fecha_fin"`
}

func (Proyecto) TableName() string {
	return "Proyecto"
}
