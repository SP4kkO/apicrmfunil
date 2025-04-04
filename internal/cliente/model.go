package cliente

import (
	"time"

	"my-crm-backend/internal/empresa"

	"gorm.io/gorm"
)

type Cliente struct {
	ID       int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome     string `json:"nome"`
	CNPJ     string `json:"cnpj"`
	Endereco string `json:"endereco"`
	Contato  string `json:"contato"`

	Empresas []empresa.Empresa `json:"empresas" gorm:"foreignKey:ClienteID"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
