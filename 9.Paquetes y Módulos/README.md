### 9. Paquetes y Módulos

En Go, los paquetes y módulos son fundamentales para organizar y reutilizar el código. Los paquetes permiten agrupar funciones, tipos y otros identificadores relacionados en una única unidad, mientras que los módulos son colecciones de paquetes que se gestionan juntos.

## Contenido

1. Introducción a los Paquetes
2. Creación y Uso de Paquetes
3. Introducción a los Módulos
4. Creación y Uso de Módulos
5. Importación de Paquetes Externos

### Introducción a los Paquetes

Un paquete en Go es una colección de archivos de Go que se encuentran en el mismo directorio y que tienen el mismo nombre de paquete al comienzo de los archivos. Los paquetes permiten dividir el código en partes más manejables y reutilizables.

#### Ejemplo de un Paquete

Supongamos que tenemos un paquete llamado `maths` con una función para sumar dos números. La estructura del directorio sería:

```
miModulo/
├── main.go
└── maths/
    └── calculator.go
```

### Creación y Uso de Paquetes

Para crear un paquete, simplemente coloca uno o más archivos de Go en un directorio y dale el mismo nombre de paquete a cada archivo. Para usar un paquete, se importa y se llama a sus funciones.

#### Ejemplo de Creación y Uso de Paquetes

`maths/maths.go`:
```go
package maths

// Sumar devuelve la suma de dos enteros
func Sumar(a, b int) int {
    return a + b
}
```

`main.go`:
```go
package main

import (
    "fmt"
    "miModulo/maths"
)

func main() {
    resultado := maths.Sumar(3, 4)
    fmt.Println("La suma es:", resultado)
}
```

**Explicación:**
- `package maths`: Define el paquete `maths`.
- `func Sumar(a, b int) int`: Declara una función `Sumar` que toma dos enteros y devuelve su suma.
- `import "miModulo/maths"`: Importa el paquete `maths` en `main.go`.
- `maths.Sumar(3, 4)`: Llama a la función `Sumar` del paquete `maths`.

### Introducción a los Módulos

Un módulo es una colección de paquetes en Go. Los módulos permiten gestionar dependencias y versionado del código. Los módulos se definen mediante un archivo `go.mod`.

### Creación y Uso de Módulos

Para crear un módulo, se usa el comando `go mod init`, que crea un archivo `go.mod` en el directorio del proyecto.

#### Ejemplo de Creación y Uso de Módulos

1. Inicializar un módulo:
```sh
go mod init miModulo
```

Esto crea un archivo `go.mod` en el directorio del proyecto.

`go.mod`:
```
module miModulo

go 1.16
```

2. Estructura del directorio:
```
miModulo/
├── go.mod
├── main.go
└── maths/
    └── calculator.go
```

3. Archivos de código:
`maths/maths.go`:
```go
package maths

// Sumar devuelve la suma de dos enteros
func Sumar(a, b int) int {
    return a + b
}
```

`main.go`:
```go
package main

import (
    "fmt"
    "miModulo/maths"
)

func main() {
    resultado := maths.Sumar(3, 4)
    fmt.Println("La suma es:", resultado)
}
```

**Explicación:**
- `go mod init miModulo`: Inicializa un módulo llamado `miModulo`.
- `module miModulo`: Define el nombre del módulo en `go.mod`.

### Importación de Paquetes Externos

Go permite importar paquetes externos para agregar funcionalidad adicional a los proyectos. Estos paquetes se gestionan mediante módulos.

#### Ejemplo de Importación de Paquetes Externos

Supongamos que queremos usar el paquete `github.com/fatih/color` para imprimir texto coloreado en la terminal.

1. Añadir el paquete externo:
```sh
go get github.com/fatih/color
```

2. Estructura del directorio:
```
miModulo/
├── go.mod
├── go.sum
├── main.go
└── maths/
    └── calculator.go
```

3. Archivos de código:
`main.go`:
```go
package main

import (
    "fmt"
    "miModulo/maths"

    "github.com/fatih/color"
)

func main() {
    resultado := maths.Sumar(3, 4)
    fmt.Println("La suma es:", resultado)

    color.Cyan("Este es un mensaje en color cian")
}
```

**Explicación:**
- `go get github.com/fatih/color`: Descarga e instala el paquete `github.com/fatih/color`.
- `import "github.com/fatih/color"`: Importa el paquete `color`.
- `color.Cyan("Este es un mensaje en color cian")`: Utiliza una función del paquete `color` para imprimir texto en color cian.

### Resumen

En este tema, hemos cubierto los conceptos fundamentales de los paquetes y módulos en Go, incluyendo cómo crear y usar paquetes, inicializar y usar módulos, y cómo importar paquetes externos. Comprender y utilizar paquetes y módulos es esencial para escribir código Go bien organizado y reutilizable, así como para gestionar dependencias de manera eficiente.