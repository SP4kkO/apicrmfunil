# Stage 1: Build
FROM golang:1.23.0-alpine AS builder

# Define o diretório de trabalho
WORKDIR /app

# Copia os arquivos de dependência e baixa as dependências
COPY go.mod go.sum ./
RUN go mod tidy

# Copia o restante do código-fonte para dentro do container
COPY . .

# Compila a aplicação, apontando para o arquivo main correto
RUN go build -o main ./cmd/server/main.go

# Stage 2: Runtime
FROM alpine:latest

WORKDIR /app

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/main .

# Exponha a porta da aplicação
EXPOSE 8080

# Comando para executar o binário
CMD ["./main"]
