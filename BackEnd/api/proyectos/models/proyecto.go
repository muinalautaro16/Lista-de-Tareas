package models

type Proyecto struct {
	ID          int    `json:"id" gorm:"column:id;primarykey"`
	Nombre      string `json:"nombre" gorm:"column:nombre"`
	Prioridad   string `json:"prioridad" gorm:"column:prioridad"`
	FechaInicio string `json:"fecha_inicio" gorm:"column:fecha_inicio"`
	FechaFin    string `json:"fecha_fin" gorm:"column:fecha_fin"`
}

func (Proyectos) TableName() string {
	return "Proyectos"
}

type Proyectos struct {
	Proyecto
}
