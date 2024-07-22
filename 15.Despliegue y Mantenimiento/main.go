### 15. Despliegue y Mantenimiento

Desplegar y mantener aplicaciones en producción es una parte crucial del ciclo de vida del desarrollo de software. En este tema, cubriremos cómo contenerizar una aplicación Go utilizando Docker, prácticas recomendadas para nombrar variables y funciones, y técnicas para compilar y desplegar aplicaciones Go.

## Contenido

1. Estándares de Nomenclatura en Go
2. Compilación y Despliegue de Aplicaciones Go
3. Contenerización con Docker
4. Prácticas de Mantenimiento

### Estándares de Nomenclatura en Go

Adherirse a los estándares de nomenclatura es importante para mantener el código legible y mantenible. Go tiene convenciones de nomenclatura claras que se deben seguir.

#### Variables y Funciones

- Utiliza `camelCase` para variables y funciones no exportadas.
- Utiliza `UpperCamelCase` para variables y funciones exportadas.

```go
package main

import "fmt"

// Variable exportada
var ExportedVariable string = "Soy exportada"

// variable no exportada
var unexportedVariable string = "No soy exportada"

// Función exportada
func ExportedFunction() {
    fmt.Println("Función exportada")
}

// función no exportada
func unexportedFunction() {
    fmt.Println("Función no exportada")
}

func main() {
    ExportedFunction()
    unexportedFunction()
}
```

#### Constantes

- Utiliza `UpperCamelCase` para constantes exportadas y `camelCase` o `ALL_CAPS` para constantes no exportadas.

```go
const Pi = 3.14         // Constante exportada
const maxRetries = 3    // Constante no exportada
const MAX_SIZE = 100    // Alternativa para constantes no exportadas
```

#### Structs y Tipos

- Utiliza `UpperCamelCase` para tipos y structs exportados.

```go
type Person struct {
    FirstName string
    LastName  string
}

type person struct {
    firstName string
    lastName  string
}
```

### Compilación y Despliegue de Aplicaciones Go

La compilación de una aplicación Go produce un ejecutable independiente. Puedes compilar y desplegar tus aplicaciones en diferentes entornos.

#### Compilación Cruzada

Go facilita la compilación cruzada para diferentes sistemas operativos y arquitecturas.

```sh
GOOS=linux GOARCH=amd64 go build -o miapp
```

**Explicación:**
- `GOOS=linux`: Especifica el sistema operativo de destino (Linux en este caso).
- `GOARCH=amd64`: Especifica la arquitectura de destino (64-bit en este caso).
- `-o miapp`: Especifica el nombre del ejecutable de salida.

### Contenerización con Docker

Docker es una herramienta popular para contenerizar aplicaciones, lo que facilita la distribución y ejecución de aplicaciones en diferentes entornos.

#### Ejemplo de Dockerfile

Crea un archivo llamado `Dockerfile` en el directorio de tu proyecto.

```Dockerfile
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
```

#### Construcción y Ejecución de la Imagen Docker

```sh
# Construir la imagen Docker
docker build -t miapp .

# Ejecutar el contenedor Docker
docker run -p 8080:8080 miapp
```

**Explicación:**
- `docker build -t miapp .`: Construye una imagen Docker etiquetada como `miapp`.
- `docker run -p 8080:8080 miapp`: Ejecuta un contenedor Docker y mapea el puerto 8080 del host al puerto 8080 del contenedor.

### Prácticas de Mantenimiento

Mantener una aplicación en producción implica monitoreo, registro de logs, y actualizaciones regulares.

#### Monitoreo y Registro de Logs

Es importante monitorear la aplicación y registrar los logs para detectar y solucionar problemas.

```go
package main

import (
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    log.Println("Solicitud recibida")
    w.Write([]byte("Hola, Mundo!"))
}

func main() {
    http.HandleFunc("/", handler)
    log.Println("Servidor escuchando en el puerto 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

**Explicación:**
- `log.Println(...)`: Registra un mensaje en los logs.
- `log.Fatal(...)`: Registra un mensaje y finaliza el programa.

#### Actualizaciones Regulares

Mantén tus dependencias actualizadas y realiza revisiones periódicas del código para mejorar la seguridad y el rendimiento.

```sh
# Actualizar dependencias
go get -u all

# Revisar el código en busca de problemas
go vet ./...
golint ./...
```

### Resumen

En este tema, hemos cubierto cómo seguir estándares de nomenclatura en Go, compilar y desplegar aplicaciones, contenerizar aplicaciones con Docker, y prácticas recomendadas para el mantenimiento de aplicaciones. Estas habilidades son esenciales para asegurar que tus aplicaciones Go sean fáciles de mantener, desplegar y escalar.