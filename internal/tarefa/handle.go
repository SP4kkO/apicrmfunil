package tarefa

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores HTTP para as operações de tarefa.
type Handler struct {
	repo Repository
}

// NovoHandler cria um novo handler para Tarefa.
func NovoHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// CriarTarefa cria uma nova tarefa.
func (h *Handler) CriarTarefa(c *gin.Context) {
	var t Tarefa
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validação dos campos obrigatórios
	if t.EmpresaID == 0 || t.Negociacao == "" || t.Assunto == "" || t.Responsavel == "" || t.Tipo == "" || t.DataAgendamento.IsZero() || t.Horario == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campos obrigatórios: EmpresaID, Negociacao, Assunto, Responsavel, Tipo, DataAgendamento, Horario"})
		return
	}
	novaTarefa, err := h.repo.Adicionar(t)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, novaTarefa)
}

// ListarTarefas retorna todas as tarefas.
func (h *Handler) ListarTarefas(c *gin.Context) {
	c.JSON(http.StatusOK, h.repo.Listar())
}

// ObterTarefa retorna uma tarefa pelo ID.
func (h *Handler) ObterTarefa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	tarefa, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tarefa)
}

// AtualizarTarefa atualiza os dados de uma tarefa existente.
func (h *Handler) AtualizarTarefa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var updated Tarefa
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// (Opcional) Você pode incluir validação de campos obrigatórios aqui
	tarefaAtualizada, err := h.repo.Atualizar(id, updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tarefaAtualizada)
}

// DeletarTarefa remove uma tarefa pelo ID.
func (h *Handler) DeletarTarefa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
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
