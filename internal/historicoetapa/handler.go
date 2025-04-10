package historicoetapa

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores HTTP para operações de histórico de etapas.
type Handler struct {
	repo Repository
}

// NovoHandler cria e retorna um novo handler para Histórico de Etapas.
func NovoHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// Criar cria um novo registro de histórico de etapa.
func (h *Handler) Criar(c *gin.Context) {
	var entrada HistoricoEtapa
	if err := c.ShouldBindJSON(&entrada); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entrada.DataAlteracao = time.Now()
	novo, err := h.repo.Adicionar(entrada)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, novo)
}

// Listar retorna todos os registros de histórico de etapas.
func (h *Handler) Listar(c *gin.Context) {
	historicos, err := h.repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, historicos)
}

// Obter retorna um registro de histórico pelo ID.
func (h *Handler) Obter(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	item, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// Atualizar modifica um registro de histórico existente.
func (h *Handler) Atualizar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var atualizado HistoricoEtapa
	if err := c.ShouldBindJSON(&atualizado); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	item, err := h.repo.Atualizar(id, atualizado)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, item)
}

// Deletar remove um registro de histórico pelo ID.
func (h *Handler) Deletar(c *gin.Context) {
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

// ListarPorNegociacao retorna os históricos de etapa filtrados pelo ID da negociação.
// Esse endpoint pode ser acessado via: GET /api/historico/historico/:negociacaoId
func (h *Handler) ListarPorNegociacao(c *gin.Context) {
	negociacaoIdStr := c.Param("negociacaoId")
	if negociacaoIdStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "negociacaoId não informado"})
		return
	}
	negociacaoId, err := strconv.Atoi(negociacaoIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "negociacaoId inválido"})
		return
	}
	itens, err := h.repo.ListarPorNegociacao(negociacaoId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, itens)
}
