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
	// Novo método para adicionar uma anotação à empresa
	AdicionarAnotacao(id int, anotacao string) (Empresa, error)
}

type repository struct {
	db *gorm.DB
}

func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Adicionar(e Empresa) (Empresa, error) {
	err := r.db.Create(&e).Error
	return e, err
}

func (r *repository) Listar() ([]Empresa, error) {
	var empresas []Empresa
	err := r.db.Find(&empresas).Error
	return empresas, err
}

func (r *repository) ObterPorID(id int) (*Empresa, error) {
	var empresa Empresa
	err := r.db.First(&empresa, id).Error
	if err != nil {
		return nil, errors.New("Empresa not found")
	}
	return &empresa, nil
}

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

func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Empresa{}, id).Error
}

// AdicionarAnotacao adiciona uma nova anotação à empresa identificada pelo ID.
func (r *repository) AdicionarAnotacao(id int, anotacao string) (Empresa, error) {
	empresa, err := r.ObterPorID(id)
	if err != nil {
		return Empresa{}, err
	}

	// Se já houver uma anotação, adiciona uma quebra de linha antes da nova

	err = r.db.Save(empresa).Error
	return *empresa, err
}
