### 13. Desarrollo Web

Go es una excelente opción para el desarrollo web debido a su eficiencia y facilidad para manejar la concurrencia. En este tema, cubriremos los conceptos básicos del desarrollo web en Go utilizando el paquete `net/http` de la biblioteca estándar, así como el uso de un framework popular llamado Gin para crear APIs RESTful.

## Contenido

1. Servidor HTTP Básico con `net/http`
2. Rutas y Manejo de Solicitudes
3. Manejo de Solicitudes y Respuestas
4. Uso del Framework Gin
5. Creación de APIs RESTful con Gin

### Servidor HTTP Básico con `net/http`

El paquete `net/http` de Go proporciona todo lo necesario para crear servidores HTTP. Un servidor HTTP básico se puede configurar en solo unas pocas líneas de código.

#### Ejemplo de Servidor HTTP Básico

```go
package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola, Mundo!")
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Servidor escuchando en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}
```

**Explicación:**
- `http.HandleFunc("/", handler)`: Asocia la ruta `/` con la función `handler`.
- `http.ListenAndServe(":8080", nil)`: Inicia el servidor en el puerto 8080.

### Rutas y Manejo de Solicitudes

Las rutas en un servidor web determinan cómo las solicitudes HTTP se manejan y responden. Se pueden definir múltiples rutas y asociarlas con diferentes manejadores.

#### Ejemplo de Definición de Rutas

```go
package main

import (
    "fmt"
    "net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Bienvenido a la página de inicio")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Esta es la página Acerca de")
}

func main() {
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/about", aboutHandler)
    fmt.Println("Servidor escuchando en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}
```

**Explicación:**
- `http.HandleFunc("/", homeHandler)`: Asocia la ruta `/` con la función `homeHandler`.
- `http.HandleFunc("/about", aboutHandler)`: Asocia la ruta `/about` con la función `aboutHandler`.

### Manejo de Solicitudes y Respuestas

Para manejar solicitudes y respuestas de manera efectiva, se utiliza el tipo `http.Request` para acceder a los detalles de la solicitud y el tipo `http.ResponseWriter` para enviar respuestas al cliente.

#### Ejemplo de Manejo de Solicitudes y Respuestas

```go
package main

import (
    "fmt"
    "net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        fmt.Fprintf(w, "Formulario recibido:\n")
        for key, value := range r.Form {
            fmt.Fprintf(w, "%s: %s\n", key, value)
        }
    } else {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}

func main() {
    http.HandleFunc("/form", formHandler)
    fmt.Println("Servidor escuchando en el puerto 8080")
    http.ListenAndServe(":8080", nil)
}
```

**Explicación:**
- `r.Method`: Verifica el método HTTP de la solicitud.
- `r.ParseForm()`: Analiza los datos del formulario.
- `r.Form`: Accede a los datos del formulario.

### Uso del Framework Gin

Gin es un framework web rápido y ligero para Go, que facilita la creación de APIs web y aplicaciones RESTful. Gin proporciona enrutamiento eficiente, manejo de middleware y otras características útiles.

#### Instalación de Gin

```sh
go get -u github.com/gin-gonic/gin
```

#### Ejemplo de Servidor Básico con Gin

```go
package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.String(200, "¡Hola, Mundo!")
    })
    r.Run(":8080") // Escucha y sirve en el puerto 8080
}
```

**Explicación:**
- `gin.Default()`: Crea un router Gin con middleware por defecto.
- `r.GET("/", func(c *gin.Context) { ... })`: Define una ruta GET para `/`.

### Creación de APIs RESTful con Gin

Crear APIs RESTful con Gin es sencillo y sigue los principios REST. A continuación se muestra un ejemplo de una API básica para manejar una lista de tareas.

#### Ejemplo de API RESTful con Gin

```go
package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

type Tarea struct {
    ID     string `json:"id"`
    Titulo string `json:"titulo"`
    Hecho  bool   `json:"hecho"`
}

var tareas = []Tarea{
    {ID: "1", Titulo: "Aprender Go", Hecho: false},
    {ID: "2", Titulo: "Construir una API", Hecho: false},
}

func obtenerTareas(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, tareas)
}

func crearTarea(c *gin.Context) {
    var nuevaTarea Tarea
    if err := c.BindJSON(&nuevaTarea); err != nil {
        return
    }
    tareas = append(tareas, nuevaTarea)
    c.IndentedJSON(http.StatusCreated, nuevaTarea)
}

func obtenerTareaPorID(c *gin.Context) {
    id := c.Param("id")
    for _, tarea := range tareas {
        if tarea.ID == id {
            c.IndentedJSON(http.StatusOK, tarea)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"mensaje": "Tarea no encontrada"})
}

func main() {
    router := gin.Default()
    router.GET("/tareas", obtenerTareas)
    router.POST("/tareas", crearTarea)
    router.GET("/tareas/:id", obtenerTareaPorID)

    router.Run(":8080")
}
```

**Explicación:**
- `router.GET("/tareas", obtenerTareas)`: Define una ruta GET para obtener todas las tareas.
- `router.POST("/tareas", crearTarea)`: Define una ruta POST para crear una nueva tarea.
- `router.GET("/tareas/:id", obtenerTareaPorID)`: Define una ruta GET para obtener una tarea por su ID.
- `c.IndentedJSON(...)`: Envía una respuesta JSON con formato indentado.

### Resumen

En este tema, hemos cubierto los conceptos básicos del desarrollo web en Go utilizando el paquete `net/http` y el framework Gin. Hemos aprendido a configurar un servidor HTTP básico, definir rutas, manejar solicitudes y respuestas, y crear APIs RESTful con Gin. Estas habilidades te permitirán desarrollar aplicaciones web eficientes y escalables en Go.