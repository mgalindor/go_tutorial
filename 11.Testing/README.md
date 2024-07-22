### 11. Testing

El testing es una parte fundamental del desarrollo de software, y Go proporciona un robusto soporte para escribir y ejecutar pruebas. En este tema, cubriremos cómo escribir pruebas unitarias, pruebas de benchmark, y cómo ejecutar y organizar estas pruebas en Go.

## Contenido

1. Introducción al Testing en Go
2. Escribir Pruebas Unitarias
3. Ejecutar Pruebas
4. Pruebas de Benchmark
5. Organización de Pruebas

### Introducción al Testing en Go

Go incluye un paquete de testing estándar llamado `testing`, que proporciona las herramientas necesarias para escribir pruebas unitarias y de benchmark. Los archivos de prueba deben seguir una convención de nomenclatura específica y ubicarse en el mismo paquete que el código que están probando.

### Escribir Pruebas Unitarias

Las pruebas unitarias en Go se escriben en archivos que terminan en `_test.go`. Cada prueba es una función que comienza con `Test` y toma un parámetro de tipo `*testing.T`.

#### Ejemplo de Pruebas Unitarias

Supongamos que tenemos una función `Sumar` en un archivo `maths.go`:

`maths.go`:
```go
package maths

func Sumar(a, b int) int {
    return a + b
}
```

Para probar esta función, creamos un archivo `maths_test.go`:

`maths_test.go`:
```go
package maths

import "testing"

func TestSumar(t *testing.T) {
    total := Sumar(3, 4)
    if total != 7 {
        t.Errorf("Sumar(3, 4) = %d; want 7", total)
    }
}
```

**Explicación:**
- `package maths`: El paquete de prueba debe coincidir con el paquete del código que se está probando.
- `func TestSumar(t *testing.T)`: Declara una función de prueba.
- `t.Errorf(...)`: Registra un error si el resultado de `Sumar(3, 4)` no es `7`.

### Ejecutar Pruebas

Para ejecutar pruebas, se utiliza el comando `go test`, que ejecuta todas las pruebas en los archivos `_test.go` del paquete actual.

```sh
go test
```

#### Ejemplo de Ejecución de Pruebas

Si ejecutamos `go test` en el directorio que contiene `maths.go` y `maths_test.go`, deberíamos ver algo como esto:

```sh
$ go test
PASS
ok      project/maths    0.002s
```

**Explicación:**
- `PASS`: Indica que todas las pruebas pasaron.
- `ok project/maths 0.002s`: Indica que las pruebas del paquete `maths` se ejecutaron correctamente en `0.002` segundos.

### Pruebas de Benchmark

Las pruebas de benchmark miden el rendimiento de funciones específicas. Se escriben en archivos `_test.go` y usan funciones que comienzan con `Benchmark` y toman un parámetro de tipo `*testing.B`.

#### Ejemplo de Pruebas de Benchmark

`maths_test.go`:
```go
package maths

import "testing"

func BenchmarkSumar(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sumar(3, 4)
    }
}
```

**Explicación:**
- `func BenchmarkSumar(b *testing.B)`: Declara una función de benchmark.
- `for i := 0; i < b.N; i++`: Ejecuta la función `Sumar` repetidamente para medir su rendimiento.

Para ejecutar pruebas de benchmark, se usa el comando `go test -bench=.`.

```sh
go test -bench=.
```

#### Ejemplo de Ejecución de Pruebas de Benchmark

```sh
$ go test -bench=.
goos: darwin
goarch: amd64
pkg: project/maths
BenchmarkSumar-8        1000000000               0.287 ns/op
PASS
ok      project/maths    2.456s
```

**Explicación:**
- `BenchmarkSumar-8 1000000000 0.287 ns/op`: Indica que `BenchmarkSumar` se ejecutó `1000000000` veces con un tiempo promedio de `0.287` nanosegundos por operación.

### Organización de Pruebas

Es importante organizar bien las pruebas para mantener el código limpio y manejable. Las pruebas deben estar en archivos separados (`*_test.go`) y ubicarse en el mismo paquete que el código que están probando.

#### Estructura de Proyecto Sugerida

```
project/
├── maths/
│   ├── maths.go
│   └── maths_test.go
├── go.mod
└── go.sum
```

### Resumen

En este tema, hemos cubierto cómo escribir y ejecutar pruebas unitarias y de benchmark en Go, y cómo organizar estas pruebas dentro de un proyecto. El testing es una parte crucial del desarrollo de software, y Go proporciona herramientas sencillas pero poderosas para asegurar que tu código sea correcto y eficiente. Dominar estas herramientas te permitirá desarrollar aplicaciones más robustas y confiables.