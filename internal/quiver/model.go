package quiver

import (
	"time"

	"gorm.io/gorm"
)

type Quiver struct {
	ID             int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome           string  `json:"nome"`
	Proposta       string  `json:"proposta"`
	Apolice        string  `json:"apolice"`
	Seguradora     string  `json:"seguradora"`
	Produto        string  `json:"produto"`
	VigenciaAgenda string  `json:"vigencia_agenda"`
	Origem         string  `json:"origem"`
	Ramo           string  `json:"ramo"`
	CpfCnpj        string  `json:"cpf_cnpj"`
	ValorPremio    float64 `json:"valor_premio"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
