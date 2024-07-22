# Etapa 1: Construcción del binario
FROM golang:1.16 AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o miapp

# Etapa 2: Crear una imagen ligera
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/miapp .
EXPOSE 8080
CMD ["./miapp"]
