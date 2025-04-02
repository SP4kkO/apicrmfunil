package empresa

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler define os manipuladores HTTP para as operações de empresa.
type Handler struct {
	repo Repository
}

// NovoHandler cria um novo handler para Empresa.
func NovoHandler(repo Repository) *Handler {
	return &Handler{repo: repo}
}

// CriarEmpresa cria uma nova empresa.
func (h *Handler) CriarEmpresa(c *gin.Context) {
	var e Empresa
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Valida campos obrigatórios: Nome e CNPJMatriz; e também ClienteDaBase
	if e.Nome == "" || e.CNPJMatriz == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Nome da empresa e CNPJ Matriz são obrigatórios"})
		return
	}
	novaEmpresa, err := h.repo.Adicionar(e)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, novaEmpresa)
}

// ListarEmpresas retorna todas as empresas.
func (h *Handler) ListarEmpresas(c *gin.Context) {
	empresas, err := h.repo.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, empresas)
}

// ObterEmpresa retorna uma empresa pelo ID.
func (h *Handler) ObterEmpresa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	empresa, err := h.repo.ObterPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, empresa)
}

// AtualizarEmpresa atualiza os dados de uma empresa existente.
func (h *Handler) AtualizarEmpresa(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var updated Empresa
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	empresaAtualizada, err := h.repo.Atualizar(id, updated)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, empresaAtualizada)
}

// DeletarEmpresa remove uma empresa pelo ID.
func (h *Handler) DeletarEmpresa(c *gin.Context) {
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
