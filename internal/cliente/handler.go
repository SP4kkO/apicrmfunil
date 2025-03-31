package cliente

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores das requisições para Cliente.
type Handler struct {
	repo *Repositorio
}

// NovoHandler cria um novo handler para cliente.
func NovoHandler(repo *Repositorio) *Handler {
	return &Handler{repo: repo}
}

// CriarCliente trata a criação de um novo cliente.
func (h *Handler) CriarCliente(c *gin.Context) {
	var novoCliente Cliente
	if err := c.ShouldBindJSON(&novoCliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clienteCriado, err := h.repo.Adicionar(novoCliente)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, clienteCriado)
}

// ListarClientes retorna todos os clientes cadastrados.
func (h *Handler) ListarClientes(c *gin.Context) {
	c.JSON(http.StatusOK, h.repo.Listar())
}

// ObterCliente busca um cliente pelo ID.
func (h *Handler) ObterCliente(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	cliente, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

// AtualizarCliente atualiza os dados de um cliente.
func (h *Handler) AtualizarCliente(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Cliente
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	clienteAtualizado, err := h.repo.Atualizar(id, updated)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clienteAtualizado)
}

// DeletarCliente remove um cliente pelo ID.
func (h *Handler) DeletarCliente(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	err = h.repo.Deletar(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
