package negociacao

import (
	"net/http"
	"strconv"
	"time"

	"my-crm-backend/internal/tarefa"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores HTTP para as operações de negociação.
type Handler struct {
	repo Repository
}

// NovoHandler cria e retorna um novo handler para negociação.
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

// AtualizarFunilHandler atualiza a etapa do funil de vendas e registra o histórico da alteração.
// Espera receber um JSON com: {"etapa_funil_vendas": "nova etapa", "alterado_por": "usuário", "observacao": "algum comentário"}
func (h *Handler) AtualizarFunilHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var payload struct {
		EtapaFunilVendas string `json:"etapa_funil_vendas"`
		AlteradoPor      string `json:"alterado_por"`
		Observacao       string `json:"observacao"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	atualizado, err := h.repo.AtualizarFunil(id, payload.EtapaFunilVendas, payload.AlteradoPor, payload.Observacao)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, atualizado)
}

// AtualizarStatusHandler atualiza o campo Status da negociação.
// Espera receber um JSON com: {"status": "novo status"}
func (h *Handler) AtualizarStatusHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var payload struct {
		Status string `json:"status"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	atualizado, err := h.repo.AtualizarStatus(id, payload.Status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, atualizado)
}

// AtualizarValoresHandler atualiza os campos ValorNegociacao e PrevisaoFechamento da negociação.
// Espera receber um JSON com: {"valor_negociacao": 123.45, "previsao_fechamento": "2025-12-31T23:59:59Z"}
func (h *Handler) AtualizarValoresHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var payload struct {
		ValorNegociacao    float64   `json:"valor_negociacao"`
		PrevisaoFechamento time.Time `json:"previsao_fechamento"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	atualizado, err := h.repo.AtualizarValores(id, payload.ValorNegociacao, payload.PrevisaoFechamento)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, atualizado)
}
