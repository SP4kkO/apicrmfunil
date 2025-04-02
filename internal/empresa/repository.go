package empresa

import (
	"errors"

	"gorm.io/gorm"
)

// Repository define as operações básicas para manipular empresas.
type Repository interface {
	Adicionar(e Empresa) (Empresa, error)
	Listar() ([]Empresa, error)
	ObterPorID(id int) (*Empresa, error)
	Atualizar(id int, updated Empresa) (Empresa, error)
	Deletar(id int) error
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria e retorna um repositório baseado em GORM.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Adicionar insere uma nova empresa no banco de dados.
func (r *repository) Adicionar(e Empresa) (Empresa, error) {
	err := r.db.Create(&e).Error
	return e, err
}

// Listar retorna todas as empresas do banco de dados.
func (r *repository) Listar() ([]Empresa, error) {
	var empresas []Empresa
	err := r.db.Find(&empresas).Error
	return empresas, err
}

// ObterPorID busca uma empresa pelo ID.
func (r *repository) ObterPorID(id int) (*Empresa, error) {
	var empresa Empresa
	err := r.db.First(&empresa, id).Error
	if err != nil {
		return nil, errors.New("Empresa not found")
	}
	return &empresa, nil
}

// Atualizar modifica os dados de uma empresa existente.
func (r *repository) Atualizar(id int, updated Empresa) (Empresa, error) {
	var empresa Empresa
	err := r.db.First(&empresa, id).Error
	if err != nil {
		return Empresa{}, errors.New("Empresa not found")
	}

	updated.ID = id
	err = r.db.Model(&empresa).Updates(updated).Error
	return updated, err
}

// Deletar remove uma empresa pelo ID.
func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Empresa{}, id).Error
}
