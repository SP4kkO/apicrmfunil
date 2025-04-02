package contato

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores HTTP para as operações de contato.
type Handler struct {
	repo Repository
}

// NovoHandler cria um novo handler para Contato.
func NovoHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// CriarContato cria um novo contato.
func (h *Handler) CriarContato(c *gin.Context) {
	var contato Contato
	if err := c.ShouldBindJSON(&contato); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if contato.Nome == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome é obrigatório"})
		return
	}
	novoContato, err := h.repo.Adicionar(contato)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, novoContato)
}

// ListarContatos retorna todos os contatos.
func (h *Handler) ListarContatos(c *gin.Context) {
	contatos, err := h.repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contatos)
}

// ObterContato retorna um contato pelo ID.
func (h *Handler) ObterContato(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	contato, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contato)
}

// AtualizarContato atualiza os dados de um contato existente.
func (h *Handler) AtualizarContato(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var updated Contato
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	contatoAtualizado, err := h.repo.Atualizar(id, updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, contatoAtualizado)
}

// DeletarContato remove um contato pelo ID.
func (h *Handler) DeletarContato(c *gin.Context) {
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
