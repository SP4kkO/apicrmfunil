package anotacao

import (
	"errors"

	"gorm.io/gorm"
)

// Repository define as operações básicas para manipulação de anotações.
type Repository interface {
	Adicionar(a Anotacao) (Anotacao, error)
	Listar() ([]Anotacao, error)
	ObterPorID(id int) (*Anotacao, error)
	Atualizar(id int, updated Anotacao) (Anotacao, error)
	Deletar(id int) error
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria uma nova instância do repositório de anotações.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Adicionar(a Anotacao) (Anotacao, error) {
	err := r.db.Create(&a).Error
	return a, err
}

func (r *repository) Listar() ([]Anotacao, error) {
	var anotacoes []Anotacao
	err := r.db.Find(&anotacoes).Error
	return anotacoes, err
}

func (r *repository) ObterPorID(id int) (*Anotacao, error) {
	var a Anotacao
	err := r.db.First(&a, id).Error
	if err != nil {
		return nil, errors.New("Anotação não encontrada")
	}
	return &a, nil
}

func (r *repository) Atualizar(id int, updated Anotacao) (Anotacao, error) {
	var a Anotacao
	err := r.db.First(&a, id).Error
	if err != nil {
		return Anotacao{}, errors.New("Anotação não encontrada")
	}
	updated.ID = id
	err = r.db.Model(&a).Updates(updated).Error
	return updated, err
}

func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Anotacao{}, id).Error
}
