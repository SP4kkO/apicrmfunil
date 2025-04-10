package negociacao

import (
	"errors"
	"time"

	"my-crm-backend/internal/historicoetapa"
	"my-crm-backend/internal/tarefa"

	"gorm.io/gorm"
)

// Repository define as operações básicas para manipulação de negociações.
type Repository interface {
	Adicionar(n Negociacao) (Negociacao, error)
	Listar() ([]Negociacao, error)
	ObterPorID(id int) (*Negociacao, error)
	Atualizar(id int, updated Negociacao) (Negociacao, error)
	Deletar(id int) error
	AdicionarTarefa(negociacaoID int, novaTarefa tarefa.Tarefa) (Negociacao, error)
	// Atualiza o funil e registra o histórico da mudança.
	AtualizarFunil(id int, novaEtapa, alteradoPor, observacao string) (Negociacao, error)
	// Métodos novos para atualização parcial:
	AtualizarStatus(id int, novoStatus string) (Negociacao, error)
	AtualizarValores(id int, valorNegociacao float64, previsaoFechamento time.Time) (Negociacao, error)
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria e retorna um repositório baseado em GORM.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Adicionar insere uma nova negociação no banco de dados.
// Se DataVencimentoApolice estiver zerada, atribui a data atual.
func (r *repository) Adicionar(n Negociacao) (Negociacao, error) {
	if n.DataVencimentoApolice.IsZero() {
		n.DataVencimentoApolice = time.Now()
	}
	err := r.db.Create(&n).Error
	return n, err
}

// Listar retorna todas as negociações com suas associações (Empresa, Contato, Tarefas e Históricos).
func (r *repository) Listar() ([]Negociacao, error) {
	var negociacoes []Negociacao
	err := r.db.
		Preload("Empresa").
		Preload("Contato").
		Preload("Tarefas").
		Preload("HistoricoEtapas").
		Find(&negociacoes).Error
	return negociacoes, err
}

// ObterPorID busca uma negociação pelo ID, incluindo as associações.
func (r *repository) ObterPorID(id int) (*Negociacao, error) {
	var negociacao Negociacao
	err := r.db.
		Preload("Empresa").
		Preload("Contato").
		Preload("Tarefas").
		Preload("HistoricoEtapas").
		First(&negociacao, id).Error
	if err != nil {
		return nil, errors.New("Negociacao not found")
	}
	return &negociacao, nil
}

// Atualizar modifica uma negociação existente.
func (r *repository) Atualizar(id int, updated Negociacao) (Negociacao, error) {
	var negociacao Negociacao
	if err := r.db.First(&negociacao, id).Error; err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}

	updated.ID = id
	err := r.db.Model(&negociacao).Updates(updated).Error
	return updated, err
}

// Deletar remove uma negociação pelo ID.
func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Negociacao{}, id).Error
}

// AdicionarTarefa adiciona uma nova tarefa à negociação especificada.
func (r *repository) AdicionarTarefa(negociacaoID int, novaTarefa tarefa.Tarefa) (Negociacao, error) {
	var negociacao Negociacao
	// Carrega a negociação com as tarefas já associadas
	err := r.db.Preload("Tarefas").First(&negociacao, negociacaoID).Error
	if err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}
	// Define a relação da nova tarefa
	novaTarefa.NegociacaoID = negociacaoID
	// Insere a nova tarefa
	if err := r.db.Create(&novaTarefa).Error; err != nil {
		return Negociacao{}, err
	}
	// Recarrega a negociação para retornar com as tarefas atualizadas
	if err := r.db.Preload("Tarefas").First(&negociacao, negociacaoID).Error; err != nil {
		return Negociacao{}, err
	}
	return negociacao, nil
}

// AtualizarFunil atualiza a etapa do funil de vendas de uma negociação e registra o histórico da alteração.
func (r *repository) AtualizarFunil(id int, novaEtapa, alteradoPor, observacao string) (Negociacao, error) {
	var negociacao Negociacao
	if err := r.db.First(&negociacao, id).Error; err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}
	oldEtapa := negociacao.EtapaFunilVendas
	// Se não houver alteração, retorna o registro atual.
	if oldEtapa == novaEtapa {
		return negociacao, nil
	}
	// Atualiza a etapa na negociação
	if err := r.db.Model(&negociacao).Update("etapa_funil_vendas", novaEtapa).Error; err != nil {
		return Negociacao{}, err
	}
	// Cria registro de histórico
	historico := historicoetapa.HistoricoEtapa{
		NegociacaoID:  negociacao.ID,
		EtapaAnterior: oldEtapa,
		EtapaAtual:    novaEtapa,
		AlteradoPor:   alteradoPor,
		Observacao:    observacao,
		DataAlteracao: time.Now(),
	}
	if err := r.db.Create(&historico).Error; err != nil {
		return Negociacao{}, err
	}
	negociacao.EtapaFunilVendas = novaEtapa
	return negociacao, nil
}

// AtualizarStatus atualiza apenas o campo Status da negociação.
func (r *repository) AtualizarStatus(id int, novoStatus string) (Negociacao, error) {
	var negociacao Negociacao
	if err := r.db.First(&negociacao, id).Error; err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}
	// Atualiza o campo "status" no banco de dados
	if err := r.db.Model(&negociacao).Update("status", novoStatus).Error; err != nil {
		return Negociacao{}, err
	}
	negociacao.Status = novoStatus
	return negociacao, nil
}

// AtualizarValores atualiza os campos ValorNegociacao e PrevisaoFechamento da negociação.
func (r *repository) AtualizarValores(id int, valorNegociacao float64, previsaoFechamento time.Time) (Negociacao, error) {
	var negociacao Negociacao
	if err := r.db.First(&negociacao, id).Error; err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}
	updates := map[string]interface{}{
		"valor_negociacao":    valorNegociacao,
		"previsao_fechamento": previsaoFechamento,
	}
	if err := r.db.Model(&negociacao).Updates(updates).Error; err != nil {
		return Negociacao{}, err
	}
	negociacao.ValorNegociacao = valorNegociacao
	negociacao.PrevisaoFechamento = previsaoFechamento
	return negociacao, nil
}
