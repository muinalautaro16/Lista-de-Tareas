package models

type Etapa struct {
	ID          int    `json:"id" gorm:"column:id;primarykey"`
	ID_proyecto int    `json:"id_proyecto" gorm:"column:id_proyecto"`
	Nombre      string `json:"nombre" gorm:"column:nombre"`
	Prioridad   string `json:"prioridad" gorm:"column:prioridad"`
}

func (Etapas) TableName() string {
	return "Etapas"
}

type Etapas struct {
	Etapa
}
