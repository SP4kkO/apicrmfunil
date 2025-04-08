package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"my-crm-backend/internal/anotacao" // import para o módulo de anotações
	"my-crm-backend/internal/cliente"
	"my-crm-backend/internal/contato"
	"my-crm-backend/internal/empresa"
	"my-crm-backend/internal/historicoetapa"
	"my-crm-backend/internal/negociacao"
	"my-crm-backend/internal/tarefa"
)

func main() {
	// Lê as variáveis de ambiente com os nomes definidos no docker-compose
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	// Define valores padrão caso alguma variável esteja vazia
	if host == "" {
		host = "db"
	}
	if user == "" {
		user = "postgres"
	}
	if dbname == "" {
		dbname = "myapp_test_db"
	}
	if password == "" {
		password = "hitallo123"
	}
	if port == "" {
		port = "5432"
	}
	if sslmode == "" {
		sslmode = "disable"
	}

	// Constrói o DSN corretamente
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	// Executa o AutoMigrate para as tabelas (incluindo o model de anotação)
	err = db.AutoMigrate(
		&cliente.Cliente{},
		&empresa.Empresa{},
		&contato.Contato{},
		&negociacao.Negociacao{},
		&tarefa.Tarefa{}, // Inclua a migração do model Tarefa
		&anotacao.Anotacao{},
	)

	if err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}

	// Inicializa o servidor Gin
	r := gin.Default()

	// Configura CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Injeta o DB nos repositórios e handlers dos módulos que usam GORM
	clienteRepo := cliente.NovoRepositorio(db)
	clienteHandler := cliente.NovoHandler(clienteRepo)

	contatoRepo := contato.NovoRepositorio(db)
	contatoHandler := contato.NovoHandler(contatoRepo)

	empresaRepo := empresa.NovoRepositorio(db)
	empresaHandler := empresa.NovoHandler(empresaRepo)

	negociacaoRepo := negociacao.NovoRepositorio(db)
	negociacaoHandler := negociacao.NovoHandler(negociacaoRepo)

	// Para Tarefa, utilizamos um repositório in-memory (conforme seu código)
	tarefaRepo := tarefa.NovoRepositorio(db)

	tarefaHandler := tarefa.NovoHandler(tarefaRepo)

	historicoRepo := historicoetapa.NovoRepositorio(db)
	historicoHandler := historicoetapa.NovoHandler(historicoRepo)

	// Injeta o DB no repositório e handler de Anotação
	anotacaoRepo := anotacao.NovoRepositorio(db)
	anotacaoHandler := anotacao.NovoHandler(anotacaoRepo)

	// Define as rotas da API
	api := r.Group("/api")
	{
		// Rotas para Clientes
		api.POST("/clientes", clienteHandler.CriarCliente)
		api.GET("/clientes", clienteHandler.ListarClientes)
		api.GET("/clientes/:id", clienteHandler.ObterCliente)
		api.PUT("/clientes/:id", clienteHandler.AtualizarCliente)
		api.DELETE("/clientes/:id", clienteHandler.DeletarCliente)

		// Rotas para Contatos
		api.POST("/contatos", contatoHandler.CriarContato)
		api.GET("/contatos", contatoHandler.ListarContatos)
		api.GET("/contatos/:id", contatoHandler.ObterContato)
		api.PUT("/contatos/:id", contatoHandler.AtualizarContato)
		api.DELETE("/contatos/:id", contatoHandler.DeletarContato)

		// Rotas para Empresas
		api.POST("/empresas", empresaHandler.CriarEmpresa)
		api.GET("/empresas", empresaHandler.ListarEmpresas)
		api.GET("/empresas/:id", empresaHandler.ObterEmpresa)
		api.PUT("/empresas/:id", empresaHandler.AtualizarEmpresa)
		api.DELETE("/empresas/:id", empresaHandler.DeletarEmpresa)
		api.POST("/empresas/:id/anotacoes", empresaHandler.AdicionarAnotacao)

		// Rotas para Tarefas
		api.POST("/tarefas", tarefaHandler.CriarTarefa)
		api.GET("/tarefas", tarefaHandler.ListarTarefas)
		api.GET("/tarefas/:id", tarefaHandler.ObterTarefa)
		api.PUT("/tarefas/:id", tarefaHandler.AtualizarTarefa)
		api.DELETE("/tarefas/:id", tarefaHandler.DeletarTarefa)

		// Rotas para Negociações
		negociacoes := api.Group("/negociacoes")
		{
			negociacoes.POST("", negociacaoHandler.CriarNegociacao)
			negociacoes.GET("", negociacaoHandler.ListarNegociacoes)
			negociacoes.GET("/:id", negociacaoHandler.ObterNegociacao)
			negociacoes.PUT("/:id", negociacaoHandler.AtualizarNegociacao)
			negociacoes.DELETE("/:id", negociacaoHandler.DeletarNegociacao)
			negociacoes.PUT("/:id/funil", negociacaoHandler.AtualizarFunilHandler)
			negociacoes.GET("/:id/historico-etapas", historicoHandler.ListarPorNegociacao)
		}

		// Rotas para Anotações
		api.POST("/anotacoes", anotacaoHandler.CriarAnotacao)
		api.GET("/anotacoes", anotacaoHandler.ListarAnotacoes)
		api.GET("/anotacoes/:id", anotacaoHandler.ObterAnotacao)
		api.PUT("/anotacoes/:id", anotacaoHandler.AtualizarAnotacao)
		api.DELETE("/anotacoes/:id", anotacaoHandler.DeletarAnotacao)
	}

	// Inicia o servidor na porta 8080
	r.Run(":8080")
}
