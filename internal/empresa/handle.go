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

// CriarEmpresa insere uma nova empresa no banco de dados.
func (h *Handler) CriarEmpresa(c *gin.Context) {
	var e Empresa
	if err := c.ShouldBindJSON(&e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Validação dos campos obrigatórios: Nome e CNPJMatriz
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

// ObterEmpresa busca uma empresa pelo ID.
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

// AtualizarEmpresa modifica os dados de uma empresa existente.
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

// AdicionarAnotacao adiciona uma nova anotação à empresa.
// Neste caso, a anotação é uma string simples. Se já existir alguma anotação,
// a nova anotação é concatenada com uma quebra de linha.
func (h *Handler) AdicionarAnotacao(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	// Espera receber um JSON no formato: {"anotacao": "texto da anotação"}
	var payload struct {
		Anotacao string `json:"anotacao"`
	}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	empresa, err := h.repo.AdicionarAnotacao(id, payload.Anotacao)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, empresa)
}
