package negociacao

import (
	"time"

	"my-crm-backend/internal/contato"
	"my-crm-backend/internal/empresa"
	"my-crm-backend/internal/historicoetapa"
	"my-crm-backend/internal/tarefa"

	"gorm.io/gorm"
)

// Negociacao representa uma negociação, vinculada a uma empresa e a um contato,
// podendo possuir diversas tarefas associadas.
type Negociacao struct {
	ID                    int             `json:"id" gorm:"primaryKey;autoIncrement"`
	EmpresaID             int             `json:"empresa_id"`
	Empresa               empresa.Empresa `json:"empresa" gorm:"foreignKey:EmpresaID"`
	ContatoID             int             `json:"contato_id"`
	Contato               contato.Contato `json:"contato"`
	NomeNegociacao        string          `json:"nome_negociacao"`
	FunilVendas           string          `json:"funil_vendas"`
	EtapaFunilVendas      string          `json:"etapa_funil_vendas"`
	Fonte                 string          `json:"fonte"`
	Campanha              string          `json:"campanha"`
	SeguradoraAtual       string          `json:"seguradora_atual"`
	DataVencimentoApolice time.Time       `json:"data_vencimento_apolice"`
	Tarefas               []tarefa.Tarefa `json:"tarefas" gorm:"foreignKey:NegociacaoID"`

	// Opcional: carregar os históricos de mudança de etapa
	HistoricoEtapas []historicoetapa.HistoricoEtapa `json:"historico_etapas,omitempty" gorm:"foreignKey:NegociacaoID"`

	// --- Campos novos adicionados ---
	Status             string    `json:"status"`              // Campo único para status
	ValorNegociacao    float64   `json:"valor_negociacao"`    // Valor da negociação
	PrevisaoFechamento time.Time `json:"previsao_fechamento"` // Data prevista para fechamento

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
