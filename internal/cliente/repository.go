package cliente

import (
	"errors"

	"gorm.io/gorm"
)

type Repositorio struct {
	db *gorm.DB
}

func NovoRepositorio(db *gorm.DB) *Repositorio {
	return &Repositorio{db: db}
}

// Adicionar insere um novo cliente no banco de dados.
func (r *Repositorio) Adicionar(c Cliente) (Cliente, error) {
	var existente Cliente
	r.db.Where("cnpj = ?", c.CNPJ).First(&existente)
	if existente.ID != 0 {
		return Cliente{}, errors.New("CNPJ already exists")
	}

	err := r.db.Create(&c).Error
	return c, err
}

// Listar retorna todos os clientes do banco de dados.
func (r *Repositorio) Listar() ([]Cliente, error) {
	var clientes []Cliente
	err := r.db.Find(&clientes).Error
	return clientes, err
}

// ObterPorID busca um cliente pelo ID.
func (r *Repositorio) ObterPorID(id int) (*Cliente, error) {
	var cliente Cliente
	err := r.db.First(&cliente, id).Error
	if err != nil {
		return nil, errors.New("Cliente not found")
	}
	return &cliente, nil
}

// Atualizar altera os dados de um cliente pelo ID.
func (r *Repositorio) Atualizar(id int, updated Cliente) (Cliente, error) {
	var cliente Cliente
	err := r.db.First(&cliente, id).Error
	if err != nil {
		return Cliente{}, errors.New("Cliente not found")
	}

	// Verifica se o novo CNPJ já existe em outro cliente.
	var existente Cliente
	r.db.Where("cnpj = ? AND id != ?", updated.CNPJ, id).First(&existente)
	if existente.ID != 0 {
		return Cliente{}, errors.New("CNPJ already exists")
	}

	updated.ID = id
	err = r.db.Model(&cliente).Updates(updated).Error
	return updated, err
}

// Deletar remove um cliente pelo ID.
func (r *Repositorio) Deletar(id int) error {
	return r.db.Delete(&Cliente{}, id).Error
}
