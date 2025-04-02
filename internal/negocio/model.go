package negocio

// Negocio representa um registro de oportunidade no CRM.
type Negocio struct {
	ID        int    `json:"id"`
	EmpresaID int    `json:"empresa_id"` // FK para a tabela de Empresas
	CNPJ      string `json:"cnpj"`
	Endereco  string `json:"endereco"`
	Contato   string `json:"contato"`
	ClienteID int    `json:"cliente_id"` // Pode referenciar um Cliente cadastrado
	Status    string `json:"status"`     // Status do funil de vendas
	Tarefa    string `json:"tarefa"`     // Tarefa a ser realizada
}

// FunilOpcoes contém os status válidos para o funil de vendas.
var FunilOpcoes = []string{
	"Lead mapeado",
	"Lead com data de retomada",
	"Visita/Reunião",
	"Adgo info QAR",
	"Em cotacao",
	"Proposta",
	"Reuniao de fechamento",
	"Assinatura de proposta",
	"Pedido permitido",
	"Pedidos 2025",
	"Pedidos mapeados",
}
