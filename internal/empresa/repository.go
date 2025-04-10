package empresa

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"my-crm-backend/internal/anotacao"
)

// Repository define as operações básicas para manipular empresas.
type Repository interface {
	Adicionar(e Empresa) (Empresa, error)
	Listar() ([]Empresa, error)
	ObterPorID(id int) (*Empresa, error)
	Atualizar(id int, updated Empresa) (Empresa, error)
	Deletar(id int) error
	AdicionarAnotacao(id int, anotacaoText string) (Empresa, error)
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria e retorna um novo repositório.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Adicionar(e Empresa) (Empresa, error) {
	err := r.db.Create(&e).Error
	return e, err
}

func (r *repository) Listar() ([]Empresa, error) {
	var empresas []Empresa
	err := r.db.
		Preload("Anotacoes").
		// Removi o Preload("Negociacoes") pois essa associação não está definida no model.
		Find(&empresas).Error
	return empresas, err
}

func (r *repository) ObterPorID(id int) (*Empresa, error) {
	var empresa Empresa
	err := r.db.
		Preload("Anotacoes").
		// Se houver relacionamento com Negociacoes, acrescente: Preload("Negociacoes").
		First(&empresa, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("empresa not found")
	} else if err != nil {
		return nil, err
	}
	return &empresa, nil
}

func (r *repository) Atualizar(id int, updated Empresa) (Empresa, error) {
	var empresa Empresa
	if err := r.db.First(&empresa, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return Empresa{}, errors.New("empresa not found")
		}
		return Empresa{}, err
	}

	updated.ID = id
	err := r.db.Model(&empresa).Updates(updated).Error
	return updated, err
}

func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Empresa{}, id).Error
}

// AdicionarAnotacao adiciona uma nova anotação à empresa identificada pelo id.
// Aqui, a operação é realizada dentro de uma transação para garantir a consistência.
func (r *repository) AdicionarAnotacao(id int, anotacaoText string) (Empresa, error) {
	var empresa Empresa

	// Executa a operação em uma transação.
	err := r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.First(&empresa, id).Error; err != nil {
			return err
		}

		// Cria uma nova anotação. Caso precise receber o "assunto" por parâmetro, ajuste aqui.
		novaAnotacao := anotacao.Anotacao{
			Data:      time.Now(),
			Assunto:   "Nova anotação", // Valor fixo; pode ser parametrizado se necessário.
			Anotacao:  anotacaoText,
			EmpresaID: empresa.ID,
		}

		if err := tx.Create(&novaAnotacao).Error; err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return Empresa{}, err
	}

	// Recarrega a empresa com as associações atualizadas.
	if err := r.db.
		Preload("Anotacoes").
		First(&empresa, id).Error; err != nil {
		return Empresa{}, errors.New("empresa not found after updating anotacoes")
	}

	return empresa, nil
}
