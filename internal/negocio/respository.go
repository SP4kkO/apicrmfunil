package negocio

import "errors"

// Repositorio simula um armazenamento in-memory para os negócios.
type Repositorio struct {
	negocios []Negocio
}

// NovoRepositorio cria e retorna um novo repositório para negócios.
func NovoRepositorio() *Repositorio {
	return &Repositorio{
		negocios: []Negocio{},
	}
}

// Adicionar insere um novo negócio.
func (r *Repositorio) Adicionar(n Negocio) (Negocio, error) {
	n.ID = len(r.negocios) + 1
	r.negocios = append(r.negocios, n)
	return n, nil
}

// Listar retorna todos os negócios.
func (r *Repositorio) Listar() []Negocio {
	return r.negocios
}

// ObterPorID busca um negócio pelo ID.
func (r *Repositorio) ObterPorID(id int) (*Negocio, error) {
	for i, n := range r.negocios {
		if n.ID == id {
			return &r.negocios[i], nil
		}
	}
	return nil, errors.New("Negocio not found")
}

// Atualizar modifica um negócio existente.
func (r *Repositorio) Atualizar(id int, updated Negocio) (Negocio, error) {
	for i, n := range r.negocios {
		if n.ID == id {
			updated.ID = id
			r.negocios[i] = updated
			return updated, nil
		}
	}
	return Negocio{}, errors.New("Negocio not found")
}

// Deletar remove um negócio pelo ID.
func (r *Repositorio) Deletar(id int) error {
	for i, n := range r.negocios {
		if n.ID == id {
			r.negocios = append(r.negocios[:i], r.negocios[i+1:]...)
			return nil
		}
	}
	return errors.New("Negocio not found")
}
