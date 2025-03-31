package main

import (
	"my-crm-backend/internal/cliente"
	"my-crm-backend/internal/negocio"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Cria a instância do Gin
	r := gin.Default()

	// Habilita o CORS para todas as rotas com configuração padrão personalizada
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Inicializa os repositórios e handlers para Cliente e Negocio
	clienteRepo := cliente.NovoRepositorio()
	clienteHandler := cliente.NovoHandler(clienteRepo)

	negocioRepo := negocio.NovoRepositorio()
	negocioHandler := negocio.NovoHandler(negocioRepo)

	// Define o grupo de rotas da API
	api := r.Group("/api")
	{
		// Rotas para a entidade Cliente
		api.POST("/clientes", clienteHandler.CriarCliente)
		api.GET("/clientes", clienteHandler.ListarClientes)
		api.GET("/clientes/:id", clienteHandler.ObterCliente)
		api.PUT("/clientes/:id", clienteHandler.AtualizarCliente)
		api.DELETE("/clientes/:id", clienteHandler.DeletarCliente)

		// Rotas para a entidade Negocio
		api.POST("/negocios", negocioHandler.CriarNegocio)
		api.GET("/negocios", negocioHandler.ListarNegocios)
		api.GET("/negocios/:id", negocioHandler.ObterNegocio)
		api.PUT("/negocios/:id", negocioHandler.AtualizarNegocio)
		api.DELETE("/negocios/:id", negocioHandler.DeletarNegocio)
	}

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
