package models

type Tarea struct {
	ID          int    `json:"id" gorm:"column:id;primarykey; not null"`
	ID_etapa    int    `json:"id_etapa" gorm:"column:id_etapa"`
	Nombre      string `json:"nombre" gorm:"column:nombre" validate:"required"`
	Descripcion string `json:"descripcion" gorm:"column:descripcion"`
	Plazo       int    `json:"plazo" gorm:"column:plazo" validate:"gt=0, lte=30, required"`
}
