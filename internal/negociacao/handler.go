package negociacao

import (
	"net/http"
	"strconv"

	"my-crm-backend/internal/tarefa"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	repo Repository
}

// NovoHandler cria um novo handler para negociação.
func NovoHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// CriarNegociacao cria uma nova negociação.
func (h *Handler) CriarNegociacao(c *gin.Context) {
	var n Negociacao
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	negociacaoCriada, err := h.repo.Adicionar(n)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, negociacaoCriada)
}

// ListarNegociacoes retorna todas as negociações.
func (h *Handler) ListarNegociacoes(c *gin.Context) {
	negociacoes, err := h.repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, negociacoes)
}

// ObterNegociacao retorna uma negociação pelo ID.
func (h *Handler) ObterNegociacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	negociacao, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, negociacao)
}

// AtualizarNegociacao atualiza uma negociação existente.
func (h *Handler) AtualizarNegociacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var updated Negociacao
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	negociacaoAtualizada, err := h.repo.Atualizar(id, updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, negociacaoAtualizada)
}

// DeletarNegociacao remove uma negociação pelo ID.
func (h *Handler) DeletarNegociacao(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	if err := h.repo.Deletar(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// AdicionarTarefaHandler adiciona uma nova tarefa a uma negociação.
func (h *Handler) AdicionarTarefaHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id")) // ID da Negociação
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var novaTarefa tarefa.Tarefa
	if err := c.ShouldBindJSON(&novaTarefa); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	negociacaoAtualizada, err := h.repo.AdicionarTarefa(id, novaTarefa)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, negociacaoAtualizada)
}

// AtualizarFunilHandler atualiza a etapa do funil de vendas de uma negociação.
func (h *Handler) AtualizarFunilHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var payload struct {
		EtapaFunilVendas string `json:"etapa_funil_vendas"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	atualizado, err := h.repo.AtualizarFunil(id, payload.EtapaFunilVendas)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, atualizado)
}
