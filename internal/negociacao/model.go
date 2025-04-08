package negociacao

import (
	"time"

	"my-crm-backend/internal/contato"
	"my-crm-backend/internal/empresa"
	"my-crm-backend/internal/tarefa"

	"gorm.io/gorm"
)

type Negociacao struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrement"`
	EmpresaID int             `json:"empresa_id"`
	Empresa   empresa.Empresa `json:"empresa" gorm:"foreignKey:EmpresaID"`
	ContatoID int             `json:"contato_id"`
	Contato   contato.Contato `json:"contato"` // para exibição

	NomeNegociacao        string          `json:"nome_negociacao"`
	FunilVendas           string          `json:"funil_vendas"`
	EtapaFunilVendas      string          `json:"etapa_funil_vendas"`
	Fonte                 string          `json:"fonte"`
	Campanha              string          `json:"campanha"`
	SeguradoraAtual       string          `json:"seguradora_atual"`
	DataVencimentoApolice time.Time       `json:"data_vencimento_apolice"`
	Tarefas               []tarefa.Tarefa `json:"tarefas" gorm:"foreignKey:NegociacaoID"` // Associação 1:N

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
