### 12. Herramientas del Lenguaje

Go proporciona una serie de herramientas útiles para mejorar la calidad del código, automatizar tareas comunes y documentar el código. Estas herramientas están integradas en el entorno de desarrollo de Go y son fáciles de usar.

## Contenido

1. `go fmt`
2. `go vet`
3. `golint`
4. `go doc`
5. `go test`
6. `go build` y `go run`
7. `go mod`

### `go fmt`

`go fmt` es una herramienta que formatea automáticamente el código fuente de Go según los estándares de estilo del lenguaje. Mantener un estilo de código consistente es importante para la legibilidad y mantenibilidad del código.

#### Ejemplo de Uso de `go fmt`

```sh
go fmt ./...
```

**Explicación:**
- `go fmt ./...`: Formatea todos los archivos Go en el directorio actual y sus subdirectorios.

### `go vet`

`go vet` analiza el código fuente de Go y señala constructos sospechosos o errores comunes. Es una herramienta estática que ayuda a identificar problemas que pueden no ser capturados por el compilador.

#### Ejemplo de Uso de `go vet`

```sh
go vet ./...
```

**Explicación:**
- `go vet ./...`: Analiza todos los archivos Go en el directorio actual y sus subdirectorios en busca de constructos sospechosos.

### `golint`

`golint` es una herramienta que sugiere mejoras en el estilo de código según las convenciones de Go. Aunque no es una herramienta oficial, es ampliamente utilizada en la comunidad de Go.

#### Instalación y Uso de `golint`

```sh
go install golang.org/x/lint/golint@latest
golint ./...
```

**Explicación:**
- `go install golang.org/x/lint/golint@latest`: Instala `golint`.
- `golint ./...`: Analiza todos los archivos Go en el directorio actual y sus subdirectorios en busca de mejoras de estilo.

### `go doc`

`go doc` muestra la documentación de los paquetes, tipos, funciones y variables en Go. Es útil para entender cómo utilizar las diferentes partes del código y las bibliotecas estándar.

#### Ejemplo de Uso de `go doc`

```sh
go doc fmt.Println
```

**Explicación:**
- `go doc fmt.Println`: Muestra la documentación de la función `Println` del paquete `fmt`.

### `go test`

`go test` ejecuta las pruebas definidas en los archivos `_test.go`. Es la herramienta principal para escribir y ejecutar pruebas unitarias en Go.

#### Ejemplo de Uso de `go test`

```sh
go test ./...
```

**Explicación:**
- `go test ./...`: Ejecuta todas las pruebas en el directorio actual y sus subdirectorios.

### `go build` y `go run`

`go build` compila el código fuente de Go en un ejecutable. `go run` compila y ejecuta el código fuente en un solo paso.

#### Ejemplo de Uso de `go build` y `go run`

```sh
go build -o miapp main.go
./miapp

go run main.go
```

**Explicación:**
- `go build -o miapp main.go`: Compila el archivo `main.go` en un ejecutable llamado `miapp`.
- `./miapp`: Ejecuta el ejecutable generado.
- `go run main.go`: Compila y ejecuta el archivo `main.go` en un solo paso.

### `go mod`

`go mod` gestiona los módulos de Go, que son colecciones de paquetes. Los módulos permiten manejar dependencias y versionado de manera eficiente.

#### Ejemplo de Uso de `go mod`

1. Inicializar un módulo:
    ```sh
    go mod init miapp
    ```

2. Añadir una dependencia:
    ```sh
    go get github.com/gin-gonic/gin
    ```

3. Actualizar las dependencias:
    ```sh
    go mod tidy
    ```

**Explicación:**
- `go mod init miapp`: Inicializa un módulo llamado `miapp`.
- `go get github.com/gin-gonic/gin`: Añade la dependencia del paquete `gin-gonic/gin`.
- `go mod tidy`: Limpia las dependencias, eliminando las que no se usan y añadiendo las que faltan.

### Resumen

En este tema, hemos cubierto varias herramientas útiles en Go que ayudan a mejorar la calidad del código, automatizar tareas comunes y gestionar dependencias. El uso de estas herramientas puede facilitar el desarrollo y mantenimiento de aplicaciones en Go, asegurando que el código sea limpio, eficiente y bien documentado.