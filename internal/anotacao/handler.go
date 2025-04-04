package anotacao

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler contém as dependências do HTTP para o módulo de anotação.
type Handler struct {
	Repo Repository
}

// NovoHandler cria uma nova instância do handler de anotação.
func NovoHandler(repo Repository) *Handler {
	return &Handler{Repo: repo}
}

// CriarAnotacao trata a requisição para criar uma nova anotação.
func (h *Handler) CriarAnotacao(c *gin.Context) {
	var a Anotacao
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Define a data atual se não for enviada
	if a.Data.IsZero() {
		a.Data = time.Now()
	}

	created, err := h.Repo.Adicionar(a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

// ListarAnotacoes trata a requisição para listar todas as anotações.
func (h *Handler) ListarAnotacoes(c *gin.Context) {
	anotacoes, err := h.Repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, anotacoes)
}

// ObterAnotacao trata a requisição para obter uma anotação por ID.
func (h *Handler) ObterAnotacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	a, err := h.Repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, a)
}

// AtualizarAnotacao trata a requisição para atualizar uma anotação existente.
func (h *Handler) AtualizarAnotacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var a Anotacao
	if err := c.ShouldBindJSON(&a); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := h.Repo.Atualizar(id, a)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeletarAnotacao trata a requisição para deletar uma anotação por ID.
func (h *Handler) DeletarAnotacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err = h.Repo.Deletar(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
