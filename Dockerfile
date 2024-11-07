# Etapa de construcción

FROM golang:1.20-alpine AS builder

# Configurar el directorio de trabajo dentro del contenedor

WORKDIR /app

# Copiar los archivos de go.mod y go.sum

COPY go.mod go.sum ./

# Descargar dependencias

RUN go mod download

# Copiar el resto de los archivos de la aplicación

COPY . .

# Construir la aplicación

RUN go build -o microblogging cmd/main.go

# Etapa final

FROM alpine:latest

# Crear una carpeta en el contenedor para la aplicación

WORKDIR /root/

# Copiar el binario de la aplicación desde la etapa de construcción

COPY --from=builder /app/microblogging .

# Configurar el puerto expuesto

EXPOSE 8080

# Ejecutar la aplicación

CMD ["./microblogging"]
