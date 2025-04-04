package contato

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Contato struct {
	ID                    int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome                  string         `json:"nome"`
	Cargo                 string         `json:"cargo,omitempty"`
	Telefones             datatypes.JSON `json:"telefones,omitempty"` // Exemplo: ["+5511999998888", "+551188887777"]
	Email                 string         `json:"email,omitempty"`
	Empresa               string         `json:"empresa,omitempty"`
	InformacoesAdicionais string         `json:"informacoes_adicionais,omitempty"`
	LinkedIn              string         `json:"linkedin,omitempty"`
	CamposPersonalizados  datatypes.JSON `json:"campos_personalizados,omitempty"`
	EDecisor              bool           `json:"e_decisor"`

	// Campo auxiliar para armazenar os IDs das negociações em que o contato está envolvido
	NegociacaoIDs []int `json:"negociacao_ids,omitempty" gorm:"-"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
