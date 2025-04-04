package historicoetapa

import (
	"time"

	"gorm.io/gorm"
)

type HistoricoEtapa struct {
	ID            uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	NegociacaoID  int       `json:"negociacao_id"`
	EtapaAnterior string    `json:"etapa_anterior"`
	EtapaAtual    string    `json:"etapa_atual"`
	AlteradoPor   string    `json:"alterado_por"`
	Observacao    string    `json:"observacao,omitempty"`
	DataAlteracao time.Time `json:"data_alteracao"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}
