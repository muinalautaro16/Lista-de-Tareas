package models

/*
Etapa Modelo general
swagger:model Etapa
*/
type Etapa struct {
	ID         int    `json:"id" gorm:"column:id;primarykey"`
	IDProyecto int    `json:"id_proyecto" gorm:"column:id_proyecto"`
	Nombre     string `json:"nombre" gorm:"column:nombre" validate:"required"`
	Estado     bool   `json:"estado" gorm:"column:estado"`
}

// TableName tabla de etapas
func (Etapas) TableName() string {
	return "Etapas"
}

// Etapas modelo de etapas
type Etapas struct {
	Etapa
}
