package cliente

import "errors"

// Repositorio simula um armazenamento in-memory para os clientes.
type Repositorio struct {
	clientes []Cliente
}

// NovoRepositorio cria e retorna um novo repositório.
func NovoRepositorio() *Repositorio {
	return &Repositorio{
		clientes: []Cliente{},
	}
}

// Adicionar insere um novo cliente, garantindo que o CNPJ seja único.
func (r *Repositorio) Adicionar(c Cliente) (Cliente, error) {
	for _, cl := range r.clientes {
		if cl.CNPJ == c.CNPJ {
			return Cliente{}, errors.New("CNPJ already exists")
		}
	}
	c.ID = len(r.clientes) + 1
	r.clientes = append(r.clientes, c)
	return c, nil
}

// Listar retorna todos os clientes cadastrados.
func (r *Repositorio) Listar() []Cliente {
	return r.clientes
}

// ObterPorID busca um cliente pelo ID.
func (r *Repositorio) ObterPorID(id int) (*Cliente, error) {
	for i, cl := range r.clientes {
		if cl.ID == id {
			return &r.clientes[i], nil
		}
	}
	return nil, errors.New("Cliente not found")
}

// Atualizar altera os dados de um cliente pelo ID.
func (r *Repositorio) Atualizar(id int, updated Cliente) (Cliente, error) {
	for i, cl := range r.clientes {
		if cl.ID == id {
			// Verifica se o novo CNPJ já existe em outro cliente.
			for _, other := range r.clientes {
				if other.CNPJ == updated.CNPJ && other.ID != id {
					return Cliente{}, errors.New("CNPJ already exists")
				}
			}
			updated.ID = id
			r.clientes[i] = updated
			return updated, nil
		}
	}
	return Cliente{}, errors.New("Cliente not found")
}

// Deletar remove um cliente pelo ID.
func (r *Repositorio) Deletar(id int) error {
	for i, cl := range r.clientes {
		if cl.ID == id {
			r.clientes = append(r.clientes[:i], r.clientes[i+1:]...)
			return nil
		}
	}
	return errors.New("Cliente not found")
}
