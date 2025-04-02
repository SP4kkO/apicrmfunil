package contato

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Contato struct {
	ID                    int            `json:"id" gorm:"primaryKey;autoIncrement"`
	ClienteID             int            `json:"cliente_id"`
	Nome                  string         `json:"nome"`
	Cargo                 string         `json:"cargo,omitempty"`
	Telefone              string         `json:"telefone,omitempty"`
	Email                 string         `json:"email,omitempty"`
	Empresa               string         `json:"empresa,omitempty"`
	InformacoesAdicionais string         `json:"informacoes_adicionais,omitempty"`
	LinkedIn              string         `json:"linkedin,omitempty"`
	CamposPersonalizados  datatypes.JSON `json:"campos_personalizados,omitempty"`
	EDecisor              bool           `json:"e_decisor"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
