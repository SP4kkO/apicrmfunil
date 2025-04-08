package negociacao

import (
	"errors"
	"time"

	"my-crm-backend/internal/tarefa"

	"gorm.io/gorm"
)

type Repository interface {
	Adicionar(n Negociacao) (Negociacao, error)
	Listar() ([]Negociacao, error)
	ObterPorID(id int) (*Negociacao, error)
	Atualizar(id int, updated Negociacao) (Negociacao, error)
	Deletar(id int) error
	AdicionarTarefa(negociacaoID int, novaTarefa tarefa.Tarefa) (Negociacao, error)
	AtualizarFunil(id int, novaEtapa string) (Negociacao, error)
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria e retorna um repositório baseado em GORM.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Adicionar insere uma nova negociação no banco de dados.
func (r *repository) Adicionar(n Negociacao) (Negociacao, error) {
	if n.DataVencimentoApolice.IsZero() {
		n.DataVencimentoApolice = time.Now()
	}
	err := r.db.Create(&n).Error
	return n, err
}

// Listar retorna todas as negociações com suas associações.
func (r *repository) Listar() ([]Negociacao, error) {
	var negociacoes []Negociacao
	err := r.db.Preload("Empresa").
		Preload("Contato").
		Preload("Tarefas").
		Find(&negociacoes).Error
	return negociacoes, err
}

// ObterPorID busca uma negociação pelo ID, incluindo suas associações.
func (r *repository) ObterPorID(id int) (*Negociacao, error) {
	var negociacao Negociacao
	err := r.db.Preload("Empresa").
		Preload("Contato").
		Preload("Tarefas").
		First(&negociacao, id).Error
	if err != nil {
		return nil, errors.New("Negociacao not found")
	}
	return &negociacao, nil
}

// Atualizar modifica uma negociação existente.
func (r *repository) Atualizar(id int, updated Negociacao) (Negociacao, error) {
	var negociacao Negociacao
	err := r.db.First(&negociacao, id).Error
	if err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}

	updated.ID = id
	err = r.db.Model(&negociacao).Updates(updated).Error
	return updated, err
}

// Deletar remove uma negociação pelo ID.
func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Negociacao{}, id).Error
}

// AdicionarTarefa adiciona uma nova tarefa à negociação especificada.
func (r *repository) AdicionarTarefa(negociacaoID int, novaTarefa tarefa.Tarefa) (Negociacao, error) {
	var negociacao Negociacao
	// Carrega a negociação com as tarefas
	err := r.db.Preload("Tarefas").First(&negociacao, negociacaoID).Error
	if err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}
	// Define a relação da nova tarefa
	novaTarefa.NegociacaoID = negociacaoID
	// Insere a nova tarefa
	err = r.db.Create(&novaTarefa).Error
	if err != nil {
		return Negociacao{}, err
	}
	// Recarrega a negociação para retornar com as tarefas atualizadas
	err = r.db.Preload("Tarefas").First(&negociacao, negociacaoID).Error
	if err != nil {
		return Negociacao{}, err
	}
	return negociacao, nil
}

// AtualizarFunil atualiza a etapa do funil de vendas de uma negociação pelo ID.
func (r *repository) AtualizarFunil(id int, novaEtapa string) (Negociacao, error) {
	var negociacao Negociacao
	err := r.db.First(&negociacao, id).Error
	if err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}

	err = r.db.Model(&negociacao).Update("etapa_funil_vendas", novaEtapa).Error
	negociacao.EtapaFunilVendas = novaEtapa
	return negociacao, err
}
