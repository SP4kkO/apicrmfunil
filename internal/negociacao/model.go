package negociacao

import (
	"time"

	"my-crm-backend/internal/contato"
	"my-crm-backend/internal/empresa"

	"gorm.io/gorm"
)

type Negociacao struct {
	ID        int             `json:"id" gorm:"primaryKey;autoIncrem-ent"`
	EmpresaID int             `json:"empresa_id"`
	Empresa   empresa.Empresa `json:"empresa" gorm:"foreignKey:EmpresaID"`
	ContatoID int             `json:"contato_id"`
	Contato   contato.Contato `json:"contato" gorm:"foreignKey:ContatoID"`

	NomeNegociacao        string    `json:"nome_negociacao"`
	FunilVendas           string    `json:"funil_vendas"`
	EtapaFunilVendas      string    `json:"etapa_funil_vendas"`
	Fonte                 string    `json:"fonte"`
	Campanha              string    `json:"campanha"`
	SeguradoraAtual       string    `json:"seguradora_atual"`
	DataVencimentoApolice time.Time `json:"data_vencimento_apolice"`
	Tarefa                string    `json:"tarefa"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
