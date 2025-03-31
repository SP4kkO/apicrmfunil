# Etapa de build
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copia arquivos de módulo e baixa as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o restante do código-fonte
COPY . .

# Compila a aplicação
RUN go build -o my-crm-backend ./cmd/server

# Etapa final
FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/my-crm-backend .
EXPOSE 8080
CMD ["./my-crm-backend"]
