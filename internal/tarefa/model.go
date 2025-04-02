package tarefa

import (
	"time"

	"my-crm-backend/internal/empresa"

	"gorm.io/gorm"
)

// Tarefa representa os dados de uma tarefa.
type Tarefa struct {
	ID                int             `json:"id" gorm:"primaryKey;autoIncrement"`
	NegociacaoID      int             `json:"negociacao_id"`
	EmpresaID         int             `json:"empresa_id"` // Nova coluna que referencia a empresa da negociação
	Empresa           empresa.Empresa `json:"empresa" gorm:"foreignKey:EmpresaID;references:ID"`
	EmpresaNegociacao string          `json:"empresa_negociacao,omitempty"` // Nome ou info da empresa, se necessário
	Negociacao        string          `json:"negociacao"`                   // Campo obrigatório que representa a negociação
	Assunto           string          `json:"assunto"`                      // Campo obrigatório
	Descricao         string          `json:"descricao,omitempty"`          // Descrição opcional
	Responsavel       string          `json:"responsavel"`                  // Campo obrigatório
	Tipo              string          `json:"tipo"`                         // Campo obrigatório
	DataAgendamento   time.Time       `json:"data_agendamento"`             // Campo obrigatório
	Horario           string          `json:"horario"`                      // Campo obrigatório (ex: "HH:MM")
	Concluida         bool            `json:"concluida"`                    // Indica se a tarefa foi concluída

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
