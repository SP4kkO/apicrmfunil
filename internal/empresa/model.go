package empresa

import (
	"time"

	"my-crm-backend/internal/anotacao"

	"gorm.io/gorm"
)

type Empresa struct {
	ID               int    `json:"id" gorm:"primaryKey;autoIncrement"`
	Nome             string `json:"nome"`
	Segmento         string `json:"segmento,omitempty"`
	URL              string `json:"url,omitempty"`
	Resumo           string `json:"resumo,omitempty"`
	TamanhoEmpresa   string `json:"tamanho_empresa,omitempty"`
	FaixaFaturamento string `json:"faixa_faturamento,omitempty"`
	CNPJMatriz       string `json:"cnpj_matriz"`
	RazaoSocial      string `json:"razao_social,omitempty"`
	TelefoneMatriz   string `json:"telefone_matriz,omitempty"`
	CEP              string `json:"cep,omitempty"`
	Cidade           string `json:"cidade,omitempty"`
	Estado           string `json:"estado,omitempty"`
	ClienteDaBase    bool   `json:"cliente_da_base"`
	ClienteID        int    `json:"cliente_id"`
	LinkedinEmpresa  string `json:"linkedin_empresa,omitempty"`
	Grupo            string `json:"grupo,omitempty"`

	// Associação com Anotações (não gera ciclo, pois anotacao não importa empresa)
	Anotacoes []anotacao.Anotacao `json:"anotacoes,omitempty" gorm:"foreignKey:EmpresaID"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
