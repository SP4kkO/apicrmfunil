package tarefa

import (
	"errors"
	"time"
)

// Repository define as operações básicas para manipular tarefas.
type Repository interface {
	Adicionar(t Tarefa) (Tarefa, error)
	Listar() []Tarefa
	ObterPorID(id int) (*Tarefa, error)
	Atualizar(id int, updated Tarefa) (Tarefa, error)
	Deletar(id int) error
}

type repository struct {
	tarefas []Tarefa
}

// NovoRepositorio cria e retorna um repositório in-memory para tarefas.
func NovoRepositorio() Repository {
	return &repository{
		tarefas: []Tarefa{},
	}
}

// Adicionar insere uma nova tarefa, atribuindo um ID automaticamente.
func (r *repository) Adicionar(t Tarefa) (Tarefa, error) {
	t.ID = len(r.tarefas) + 1
	// Se a data não for definida, atribui a data atual
	if t.DataAgendamento.IsZero() {
		t.DataAgendamento = time.Now()
	}
	r.tarefas = append(r.tarefas, t)
	return t, nil
}

// Listar retorna todas as tarefas.
func (r *repository) Listar() []Tarefa {
	return r.tarefas
}

// ObterPorID busca uma tarefa pelo ID.
func (r *repository) ObterPorID(id int) (*Tarefa, error) {
	for i, t := range r.tarefas {
		if t.ID == id {
			return &r.tarefas[i], nil
		}
	}
	return nil, errors.New("Tarefa not found")
}

// Atualizar modifica os dados de uma tarefa existente.
func (r *repository) Atualizar(id int, updated Tarefa) (Tarefa, error) {
	for i, t := range r.tarefas {
		if t.ID == id {
			updated.ID = id
			r.tarefas[i] = updated
			return updated, nil
		}
	}
	return Tarefa{}, errors.New("Tarefa not found")
}

// Deletar remove uma tarefa pelo ID.
func (r *repository) Deletar(id int) error {
	for i, t := range r.tarefas {
		if t.ID == id {
			r.tarefas = append(r.tarefas[:i], r.tarefas[i+1:]...)
			return nil
		}
	}
	return errors.New("Tarefa not found")
}
