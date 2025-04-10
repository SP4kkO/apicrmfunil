package historicoetapa

import (
	"time"

	"gorm.io/gorm"
)

// HistoricoEtapa representa o registro de mudança de etapa em uma negociação.
type HistoricoEtapa struct {
	ID            uint           `json:"id" gorm:"primaryKey;autoIncrement"`
	NegociacaoID  int            `json:"negociacao_id"`
	EtapaAnterior string         `json:"etapa_anterior"`
	EtapaAtual    string         `json:"etapa_atual"`
	AlteradoPor   string         `json:"alterado_por"`
	Observacao    string         `json:"observacao,omitempty"`
	DataAlteracao time.Time      `json:"data_alteracao"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

// TableName retorna o nome da tabela que o GORM deverá usar.
func (HistoricoEtapa) TableName() string {
	return "historico_etapas"
}
