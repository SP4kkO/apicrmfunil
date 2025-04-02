package negociacao

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

// Repository define as operações básicas para manipular negociações.
type Repository interface {
	Adicionar(n Negociacao) (Negociacao, error)
	Listar() ([]Negociacao, error)
	ObterPorID(id int) (*Negociacao, error)
	Atualizar(id int, updated Negociacao) (Negociacao, error)
	Deletar(id int) error
	AtualizarTarefa(id int, novaTarefa string) (Negociacao, error)
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

// Listar retorna todas as negociações do banco de dados.
func (r *repository) Listar() ([]Negociacao, error) {
	var negociacoes []Negociacao
	err := r.db.Preload("Empresa").Preload("Contato").Find(&negociacoes).Error
	return negociacoes, err
}

// ObterPorID busca uma negociação pelo ID.
func (r *repository) ObterPorID(id int) (*Negociacao, error) {
	var negociacao Negociacao
	err := r.db.Preload("Empresa").Preload("Contato").First(&negociacao, id).Error
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

// AtualizarTarefa atualiza apenas o campo "tarefa" de uma negociação pelo ID.
func (r *repository) AtualizarTarefa(id int, novaTarefa string) (Negociacao, error) {
	var negociacao Negociacao
	err := r.db.First(&negociacao, id).Error
	if err != nil {
		return Negociacao{}, errors.New("Negociacao not found")
	}

	err = r.db.Model(&negociacao).Update("tarefa", novaTarefa).Error
	negociacao.Tarefa = novaTarefa
	return negociacao, err
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
