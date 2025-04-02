package contato

import (
	"errors"

	"gorm.io/gorm"
)

// Repository define as operações básicas para manipular contatos.
type Repository interface {
	Adicionar(c Contato) (Contato, error)
	Listar() ([]Contato, error)
	ObterPorID(id int) (*Contato, error)
	Atualizar(id int, updated Contato) (Contato, error)
	Deletar(id int) error
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria e retorna um repositório baseado em GORM.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Adicionar insere um novo contato no banco de dados.
func (r *repository) Adicionar(c Contato) (Contato, error) {
	err := r.db.Create(&c).Error
	return c, err
}

// Listar retorna todos os contatos do banco de dados.
func (r *repository) Listar() ([]Contato, error) {
	var contatos []Contato
	err := r.db.Find(&contatos).Error
	return contatos, err
}

// ObterPorID busca um contato pelo ID.
func (r *repository) ObterPorID(id int) (*Contato, error) {
	var contato Contato
	err := r.db.First(&contato, id).Error
	if err != nil {
		return nil, errors.New("Contato not found")
	}
	return &contato, nil
}

// Atualizar modifica os dados de um contato existente.
func (r *repository) Atualizar(id int, updated Contato) (Contato, error) {
	var contato Contato
	err := r.db.First(&contato, id).Error
	if err != nil {
		return Contato{}, errors.New("Contato not found")
	}

	updated.ID = id
	err = r.db.Model(&contato).Updates(updated).Error
	return updated, err
}

// Deletar remove um contato pelo ID.
func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Contato{}, id).Error
}
