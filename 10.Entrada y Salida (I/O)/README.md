### 10. Entrada y Salida (I/O)

En Go, las operaciones de entrada y salida (I/O) son esenciales para interactuar con el sistema de archivos, leer y escribir datos, y manipular cadenas y JSON. En este tema, cubriremos cómo manejar archivos, trabajar con cadenas y JSON, y algunas funciones comunes de I/O.

## Contenido

1. Manejo de Archivos
2. Lectura y Escritura de Archivos
3. Manipulación de Cadenas
4. Trabajo con JSON

### Manejo de Archivos

El paquete `os` en Go proporciona funciones para manejar archivos y directorios, como abrir, crear y eliminar archivos.

#### Ejemplo de Manejo de Archivos

```go
package main

import (
    "fmt"
    "os"
)

func main() {
    // Crear un archivo
    archivo, err := os.Create("archivo.txt")
    if err != nil {
        fmt.Println("Error al crear el archivo:", err)
        return
    }
    defer archivo.Close()

    // Escribir en el archivo
    _, err = archivo.WriteString("Hola, Go!\n")
    if err != nil {
        fmt.Println("Error al escribir en el archivo:", err)
        return
    }

    fmt.Println("Archivo creado y escrito con éxito")
}
```

**Explicación:**
- `os.Create("archivo.txt")`: Crea un archivo nuevo llamado `archivo.txt`.
- `archivo.WriteString("Hola, Go!\n")`: Escribe una cadena en el archivo.
- `defer archivo.Close()`: Asegura que el archivo se cierre al final del programa.

### Lectura y Escritura de Archivos

El paquete `bufio` proporciona un búfer de entrada y salida eficiente, que se utiliza comúnmente para leer y escribir datos de archivos.

#### Ejemplo de Lectura de Archivos

```go
package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    // Abrir un archivo existente
    archivo, err := os.Open("archivo.txt")
    if err != nil {
        fmt.Println("Error al abrir el archivo:", err)
        return
    }
    defer archivo.Close()

    // Leer el archivo línea por línea
    scanner := bufio.NewScanner(archivo)
    for scanner.Scan() {
        fmt.Println("Línea:", scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error al leer el archivo:", err)
    }
}
```

**Explicación:**
- `os.Open("archivo.txt")`: Abre un archivo existente llamado `archivo.txt`.
- `bufio.NewScanner(archivo)`: Crea un nuevo scanner para leer el archivo línea por línea.
- `scanner.Scan()`: Lee la siguiente línea del archivo.
- `scanner.Text()`: Obtiene el texto de la línea actual.

### Manipulación de Cadenas

El paquete `strings` proporciona funciones para manipular cadenas de texto, como concatenar, dividir, reemplazar y buscar substrings.

#### Ejemplo de Manipulación de Cadenas

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    texto := "Hola, Go!"

    // Convertir a mayúsculas
    mayusculas := strings.ToUpper(texto)
    fmt.Println("Mayúsculas:", mayusculas)

    // Reemplazar una subcadena
    reemplazado := strings.ReplaceAll(texto, "Go", "Golang")
    fmt.Println("Reemplazado:", reemplazado)

    // Dividir una cadena
    partes := strings.Split(texto, ", ")
    fmt.Println("Partes:", partes)

    // Comprobar si contiene una subcadena
    contiene := strings.Contains(texto, "Go")
    fmt.Println("Contiene 'Go':", contiene)
}
```

**Explicación:**
- `strings.ToUpper(texto)`: Convierte la cadena `texto` a mayúsculas.
- `strings.ReplaceAll(texto, "Go", "Golang")`: Reemplaza todas las ocurrencias de `Go` por `Golang` en la cadena `texto`.
- `strings.Split(texto, ", ")`: Divide la cadena `texto` en partes utilizando `, ` como delimitador.
- `strings.Contains(texto, "Go")`: Comprueba si la cadena `texto` contiene la subcadena `Go`.

### Trabajo con JSON

El paquete `encoding/json` proporciona funciones para codificar y decodificar datos en formato JSON. Esto es útil para trabajar con datos estructurados y para la comunicación con APIs.

#### Ejemplo de Trabajo con JSON

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

type Persona struct {
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
}

func main() {
    persona := Persona{Nombre: "Juan", Edad: 30}

    // Codificar a JSON
    jsonData, err := json.Marshal(persona)
    if err != nil {
        fmt.Println("Error al codificar a JSON:", err)
        return
    }
    fmt.Println("JSON:", string(jsonData))

    // Escribir JSON a un archivo
    err = os.WriteFile("persona.json", jsonData, 0644)
    if err != nil {
        fmt.Println("Error al escribir el archivo JSON:", err)
        return
    }

    // Leer JSON desde un archivo
    data, err := os.ReadFile("persona.json")
    if err != nil {
        fmt.Println("Error al leer el archivo JSON:", err)
        return
    }

    // Decodificar desde JSON
    var personaDecodificada Persona
    err = json.Unmarshal(data, &personaDecodificada)
    if err != nil {
        fmt.Println("Error al decodificar JSON:", err)
        return
    }
    fmt.Println("Persona decodificada:", personaDecodificada)
}
```

**Explicación:**
- `json.Marshal(persona)`: Codifica la struct `persona` en formato JSON.
- `os.WriteFile("persona.json", jsonData, 0644)`: Escribe los datos JSON en un archivo.
- `os.ReadFile("persona.json")`: Lee los datos JSON desde un archivo.
- `json.Unmarshal(data, &personaDecodificada)`: Decodifica los datos JSON en una struct `Persona`.

### Resumen

En este tema, hemos cubierto los conceptos fundamentales de las operaciones de entrada y salida en Go, incluyendo cómo manejar archivos, leer y escribir datos, manipular cadenas y trabajar con JSON. Estas habilidades son esenciales para interactuar con el sistema de archivos, procesar datos y comunicarse con APIs en tus programas en Go.