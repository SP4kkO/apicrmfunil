package tarefa

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Adicionar(t Tarefa) (Tarefa, error)
	Listar() ([]Tarefa, error)
	ObterPorID(id int) (*Tarefa, error)
	Atualizar(id int, updated Tarefa) (Tarefa, error)
	Deletar(id int) error
}

type repository struct {
	db *gorm.DB
}

// NovoRepositorio cria e retorna um repositório baseado em GORM para tarefas.
func NovoRepositorio(db *gorm.DB) Repository {
	return &repository{db: db}
}

// Adicionar insere uma nova tarefa no banco de dados.
func (r *repository) Adicionar(t Tarefa) (Tarefa, error) {
	// Se a data não for definida, atribui a data atual
	if t.DataAgendamento.IsZero() {
		t.DataAgendamento = time.Now()
	}
	err := r.db.Create(&t).Error
	return t, err
}

// Listar retorna todas as tarefas.
func (r *repository) Listar() ([]Tarefa, error) {
	var tarefas []Tarefa
	err := r.db.Find(&tarefas).Error
	return tarefas, err
}

// ObterPorID busca uma tarefa pelo ID.
func (r *repository) ObterPorID(id int) (*Tarefa, error) {
	var t Tarefa
	err := r.db.First(&t, id).Error
	if err != nil {
		return nil, errors.New("Tarefa not found")
	}
	return &t, nil
}

// Atualizar modifica os dados de uma tarefa existente.
func (r *repository) Atualizar(id int, updated Tarefa) (Tarefa, error) {
	var tarefa Tarefa
	err := r.db.First(&tarefa, id).Error
	if err != nil {
		return Tarefa{}, errors.New("Tarefa not found")
	}
	updated.ID = id
	err = r.db.Model(&tarefa).Updates(updated).Error
	return updated, err
}

// Deletar remove uma tarefa pelo ID.
func (r *repository) Deletar(id int) error {
	return r.db.Delete(&Tarefa{}, id).Error
}
