FROM golang:latest
LABEL maintainer="Souhaib EM <souhaibem@dragonlab.team>"

# Crea el directorio de trabajo
RUN mkdir /app
WORKDIR /app

# Copia los archivos del proyecto al contenedor
COPY . .

# Descarga las dependencias
RUN go mod download

# Vendoriza las dependencias
RUN go mod vendor

# Compila la aplicación
RUN go build -o main .

# Expone el puerto en el que escucha la aplicación
EXPOSE 8080

# Ejecuta la aplicación
CMD ["./main"]
