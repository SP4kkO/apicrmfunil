package quiver

import (
	"errors"

	"gorm.io/gorm"
)

type Repository interface {
	Adicionar(q Quiver) (Quiver, error)
	Listar() ([]Quiver, error)
	ObterPorID(id int) (*Quiver, error)
	Atualizar(id int, q Quiver) (Quiver, error)
	Deletar(id int) error
}

type repository struct {
	db *gorm.DB
}

func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Adicionar(q Quiver) (Quiver, error) {
	err := r.db.Create(&q).Error
	return q, err
}

func (r *repository) Listar() ([]Quiver, error) {
	var quivers []Quiver
	err := r.db.Find(&quivers).Error
	return quivers, err
}

func (r *repository) ObterPorID(id int) (*Quiver, error) {
	var q Quiver
	if err := r.db.First(&q, id).Error; err != nil {
		return nil, errors.New("registro não encontrado")
	}
	return &q, nil
}

func (r *repository) Atualizar(id int, q Quiver) (Quiver, error) {
	var existente Quiver
	if err := r.db.First(&existente, id).Error; err != nil {
		return Quiver{}, errors.New("registro não encontrado")
	}
	q.ID = id
	if err := r.db.Model(&existente).Updates(q).Error; err != nil {
		return Quiver{}, err
	}
	return q, nil
}

func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Quiver{}, id).Error
}
