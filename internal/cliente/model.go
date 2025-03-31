package cliente

// Cliente representa uma empresa ou cliente no sistema de CRM.
type Cliente struct {
	ID       int    `json:"id"`
	Nome     string `json:"nome"`
	CNPJ     string `json:"cnpj"`
	Endereco string `json:"endereco"`
	Contato  string `json:"contato"`
}
