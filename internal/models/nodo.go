package models

type Nodo struct {
	ID         uint   `gorm:"primaryKey" json:"id" required:"true"`
	EleccionID uint   `gorm:"not null;index" json:"eleccion_id" required:"true"`
	Activo     bool   `gorm:"not null" json:"activo" required:"true"`
	Secreto    string `gorm:"type:text;not null" json:"-" validate:"required_if=Activo true"`
}

func (n *Nodo) Validate() error {
	return validate.Struct(n)
}
