package negocio

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores para as requisições de Negocio.
type Handler struct {
	repo *Repositorio
}

// NovoHandler cria um novo handler para negócio.
func NovoHandler(repo *Repositorio) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) AtualizarStatusHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
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

// CriarNegocio trata a criação de um novo negócio.
func (h *Handler) CriarNegocio(c *gin.Context) {
	var novoNegocio Negocio
	if err := c.ShouldBindJSON(&novoNegocio); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Valida se o status informado é um dos permitidos.
	valid := false
	for _, s := range FunilOpcoes {
		if s == novoNegocio.Status {
			valid = true
			break
		}
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status de funil inválido"})
		return
	}
	negocioCriado, err := h.repo.Adicionar(novoNegocio)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, negocioCriado)
}

// ListarNegocios retorna todos os negócios.
func (h *Handler) ListarNegocios(c *gin.Context) {
	c.JSON(http.StatusOK, h.repo.Listar())
}

// ObterNegocio busca um negócio pelo ID.
func (h *Handler) ObterNegocio(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	negocio, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, negocio)
}

// AtualizarNegocio atualiza os dados de um negócio.
func (h *Handler) AtualizarNegocio(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var updated Negocio
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Valida o status atualizado.
	valid := false
	for _, s := range FunilOpcoes {
		if s == updated.Status {
			valid = true
			break
		}
	}
	if !valid {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status de funil inválido"})
		return
	}
	negocioAtualizado, err := h.repo.Atualizar(id, updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, negocioAtualizado)
}

// AtualizarTarefaHandler atualiza somente o campo "tarefa" de um negócio.
func (h *Handler) AtualizarTarefaHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var payload struct {
		Tarefa string `json:"tarefa"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	atualizado, err := h.repo.AtualizarTarefa(id, payload.Tarefa)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, atualizado)
}

// DeletarNegocio remove um negócio pelo ID.
func (h *Handler) DeletarNegocio(c *gin.Context) {
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
