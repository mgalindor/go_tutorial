### 7. Errores

El manejo de errores en Go es fundamental para escribir aplicaciones robustas y confiables. Go utiliza un enfoque explícito para el manejo de errores, en lugar de excepciones, lo que promueve un control más claro y directo sobre los flujos de error.

## Contenido

1. Introducción al Manejo de Errores
2. Defer, Panic y Recover
3. Creación de Errores Personalizados
4. Manejo de Errores Comunes

### Introducción al Manejo de Errores

En Go, los errores se manejan mediante valores de retorno adicionales. Las funciones que pueden resultar en un error típicamente devuelven un valor de tipo `error` junto con otros valores de retorno. El tipo `error` es una interfaz estándar en Go.

#### Ejemplo de Manejo de Errores

```go
package main

import (
    "errors"
    "fmt"
)

// Función que puede devolver un error
func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("no se puede dividir por cero")
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(4, 0)
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Resultado:", resultado)
    }
}
```

**Explicación:**
- `func dividir(a, b float64) (float64, error)`: Declara una función `dividir` que devuelve un valor `float64` y un `error`.
- `if b == 0 { return 0, errors.New("no se puede dividir por cero") }`: Devuelve un error si `b` es igual a 0.
- `resultado, err := dividir(4, 0)`: Llama a la función `dividir` y maneja el posible error retornado.

### Defer, Panic y Recover

Go proporciona tres mecanismos para el manejo avanzado de errores y la recuperación de situaciones de pánico: `defer`, `panic` y `recover`.

#### Defer

La palabra clave `defer` pospone la ejecución de una función hasta que la función que la contiene termine. `defer` es útil para tareas de limpieza, como cerrar archivos o liberar recursos.

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    archivo, err := os.Open("example.txt")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer archivo.Close() // Asegura que el archivo se cierre al final

    // Leer y procesar el archivo
    fmt.Println("Archivo abierto con éxito")
}
```

**Explicación:**
- `defer archivo.Close()`: Pospone el cierre del archivo hasta que `main` termine.

#### Panic

La función `panic` se utiliza para detener la ejecución normal del programa. Se usa en situaciones donde el programa no puede continuar.

```go
package main

import (
    "fmt"
)

func causarPanico() {
    panic("ocurrió un problema grave")
}

func main() {
    fmt.Println("Inicio del programa")
    causarPanico()
    fmt.Println("Este mensaje no se imprimirá")
}
```

**Explicación:**
- `panic("ocurrió un problema grave")`: Detiene la ejecución normal y entra en modo pánico.

#### Recover

La función `recover` se utiliza para recuperar de un `panic` y permitir que el programa continúe ejecutándose. `recover` solo es útil dentro de una función `defer`.

```go
package main

import (
    "fmt"
)

func causarPanico() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado del pánico:", r)
        }
    }()
    panic("ocurrió un problema grave")
}

func main() {
    fmt.Println("Inicio del programa")
    causarPanico()
    fmt.Println("El programa se recuperó del pánico")
}
```

**Explicación:**
- `defer func() { if r := recover(); r != nil { ... } }()`: Define una función diferida que recupera del pánico si ocurre.

### Creación de Errores Personalizados

Puedes crear errores personalizados para proporcionar más contexto sobre lo que salió mal. Esto se hace utilizando la función `errors.New` o el tipo `fmt.Errorf`.

#### Ejemplo de Errores Personalizados

```go
package main

import (
    "errors"
    "fmt"
)

var ErrDividirPorCero = errors.New("no se puede dividir por cero")

func dividir(a, b float64) (float64, error) {
    if b == 0 {
        return 0, ErrDividirPorCero
    }
    return a / b, nil
}

func main() {
    resultado, err := dividir(4, 0)
    if err != nil {
        if errors.Is(err, ErrDividirPorCero) {
            fmt.Println("Error específico:", err)
        } else {
            fmt.Println("Error:", err)
        }
    } else {
        fmt.Println("Resultado:", resultado)
    }
}
```

**Explicación:**
- `var ErrDividirPorCero = errors.New("no se puede dividir por cero")`: Declara un error personalizado.
- `if errors.Is(err, ErrDividirPorCero) { ... }`: Comprueba si el error es del tipo `ErrDividirPorCero`.

### Manejo de Errores Comunes

El manejo de errores comunes implica la captura y el tratamiento adecuado de errores típicos como el acceso a archivos, la conversión de tipos y la comunicación en red.

#### Ejemplo de Manejo de Errores Comunes

```go
package main

import (
    "fmt"
    "os"
)

func leerArchivo(nombre string) error {
    archivo, err := os.Open(nombre)
    if err != nil {
        return fmt.Errorf("no se pudo abrir el archivo %s: %v", nombre, err)
    }
    defer archivo.Close()

    // Procesar el archivo
    fmt.Println("Archivo leído con éxito")
    return nil
}

func main() {
    err := leerArchivo("example.txt")
    if err != nil {
        fmt.Println("Error:", err)
    }
}
```

**Explicación:**
- `archivo, err := os.Open(nombre)`: Intenta abrir un archivo.
- `return fmt.Errorf("no se pudo abrir el archivo %s: %v", nombre, err)`: Devuelve un error detallado si ocurre un problema.

### Resumen

En este tema, hemos cubierto los conceptos y técnicas esenciales para el manejo de errores en Go, incluyendo el uso de valores de retorno de error, `defer`, `panic`, y `recover`, la creación de errores personalizados y el manejo de errores comunes. Entender y aplicar estas técnicas es crucial para escribir programas robustos y confiables en Go.