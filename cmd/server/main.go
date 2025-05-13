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

	"my-crm-backend/internal/anotacao"
	"my-crm-backend/internal/cliente"
	"my-crm-backend/internal/contato"
	"my-crm-backend/internal/empresa"
	"my-crm-backend/internal/historicoetapa"
	"my-crm-backend/internal/negociacao"
	"my-crm-backend/internal/quiver"
	"my-crm-backend/internal/tarefa"
)

func main() {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

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

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar com o banco de dados: %v", err)
	}

	err = db.AutoMigrate(
		&cliente.Cliente{},
		&empresa.Empresa{},
		&contato.Contato{},
		&negociacao.Negociacao{},
		&tarefa.Tarefa{},
		&anotacao.Anotacao{},
		&historicoetapa.HistoricoEtapa{},
		&quiver.Quiver{},
	)
	if err != nil {
		log.Fatalf("Erro ao migrar o banco de dados: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	clienteRepo := cliente.NovoRepositorio(db)
	clienteHandler := cliente.NovoHandler(clienteRepo)

	contatoRepo := contato.NovoRepositorio(db)
	contatoHandler := contato.NovoHandler(contatoRepo)

	empresaRepo := empresa.NovoRepositorio(db)
	empresaHandler := empresa.NovoHandler(empresaRepo)

	negociacaoRepo := negociacao.NovoRepositorio(db)
	negociacaoHandler := negociacao.NovoHandler(negociacaoRepo)

	tarefaRepo := tarefa.NovoRepositorio(db)
	tarefaHandler := tarefa.NovoHandler(tarefaRepo)

	historicoRepo := historicoetapa.NovoRepositorio(db)
	historicoHandler := historicoetapa.NovoHandler(historicoRepo)

	anotacaoRepo := anotacao.NovoRepositorio(db)
	anotacaoHandler := anotacao.NovoHandler(anotacaoRepo)

	quiverRepo := quiver.NovoRepositorio(db)
	quiverHandler := quiver.NovoHandler(quiverRepo)

	api := r.Group("/api")
	{
		api.POST("/clientes", clienteHandler.CriarCliente)
		api.GET("/clientes", clienteHandler.ListarClientes)
		api.GET("/clientes/:id", clienteHandler.ObterCliente)
		api.PUT("/clientes/:id", clienteHandler.AtualizarCliente)
		api.DELETE("/clientes/:id", clienteHandler.DeletarCliente)

		api.POST("/contatos", contatoHandler.CriarContato)
		api.GET("/contatos", contatoHandler.ListarContatos)
		api.GET("/contatos/:id", contatoHandler.ObterContato)
		api.PUT("/contatos/:id", contatoHandler.AtualizarContato)
		api.DELETE("/contatos/:id", contatoHandler.DeletarContato)

		api.POST("/empresas", empresaHandler.CriarEmpresa)
		api.GET("/empresas", empresaHandler.ListarEmpresas)
		api.GET("/empresas/:id", empresaHandler.ObterEmpresa)
		api.PUT("/empresas/:id", empresaHandler.AtualizarEmpresa)
		api.DELETE("/empresas/:id", empresaHandler.DeletarEmpresa)
		api.POST("/empresas/:id/anotacoes", empresaHandler.AdicionarAnotacao)

		api.POST("/tarefas", tarefaHandler.CriarTarefa)
		api.GET("/tarefas", tarefaHandler.ListarTarefas)
		api.GET("/tarefas/:id", tarefaHandler.ObterTarefa)
		api.PUT("/tarefas/:id", tarefaHandler.AtualizarTarefa)
		api.DELETE("/tarefas/:id", tarefaHandler.DeletarTarefa)

		negociacoes := api.Group("/negociacoes")
		{
			negociacoes.POST("", negociacaoHandler.CriarNegociacao)
			negociacoes.GET("", negociacaoHandler.ListarNegociacoes)
			negociacoes.GET(":id", negociacaoHandler.ObterNegociacao)
			negociacoes.PUT(":id", negociacaoHandler.AtualizarNegociacao)
			negociacoes.DELETE(":id", negociacaoHandler.DeletarNegociacao)
			negociacoes.PUT(":id/funil", negociacaoHandler.AtualizarFunilHandler)
			negociacoes.PUT(":id/status", negociacaoHandler.AtualizarStatusHandler)
			negociacoes.PUT(":id/valores", negociacaoHandler.AtualizarValoresHandler)
			negociacoes.GET(":id/historico-etapas", historicoHandler.ListarPorNegociacao)
		}

		historico := api.Group("/historico")
		{
			historico.GET("/historico/:negociacaoId", historicoHandler.ListarPorNegociacao)
			historico.POST("", historicoHandler.Criar)
			historico.GET(":id", historicoHandler.Obter)
			historico.PUT(":id", historicoHandler.Atualizar)
			historico.DELETE(":id", historicoHandler.Deletar)
		}

		api.POST("/anotacoes", anotacaoHandler.CriarAnotacao)
		api.GET("/anotacoes", anotacaoHandler.ListarAnotacoes)
		api.GET("/anotacoes/:id", anotacaoHandler.ObterAnotacao)
		api.PUT("/anotacoes/:id", anotacaoHandler.AtualizarAnotacao)
		api.DELETE("/anotacoes/:id", anotacaoHandler.DeletarAnotacao)

		// Rotas para Quiver
		quivers := api.Group("/quivers")
		{
			quivers.POST("", quiverHandler.Criar)
			quivers.GET("", quiverHandler.Listar)
			quivers.GET(":id", quiverHandler.ObterPorID)
			quivers.PUT(":id", quiverHandler.Atualizar)
			quivers.DELETE(":id", quiverHandler.Deletar)
		}
	}

	r.Run(":8082")
}
