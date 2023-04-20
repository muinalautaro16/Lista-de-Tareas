package my

import (
	"ASTRIC/BackEnd/ASTRIC/global"
	"time"
)

// Actor Modelo general
/*
swagger:model Actor
*/
type Actor struct {
	//swagger:ignore
	ID          int       `json:"id,omitempty" gorm:"column:id;primaryKey"`
	CreatedAt   time.Time `json:"created_at,omitempty" gorm:"column:created_at"`
	UpdatedAt   time.Time `json:"updated_at,omitempty" gorm:"column:updated_at"`
	Descripcion string    `json:"descripcion,omitempty" gorm:"column:descripcion"`
}

// TableName sobreescritura de tabla
func (Actor) TableName() string {
	tabla := "actors"
	global.SetTablaResponceWS(tabla)
	return tabla
}
