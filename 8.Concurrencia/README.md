### 8. Concurrencia

Go tiene soporte incorporado para la concurrencia, lo que facilita la escritura de programas que realizan múltiples tareas de manera simultánea. La concurrencia en Go se maneja principalmente a través de goroutines y canales.

## Contenido

1. Introducción a la Concurrencia
2. Goroutines
3. Canales
4. Select
5. Sincronización con `sync.WaitGroup`
6. Mutexes

### Introducción a la Concurrencia

La concurrencia permite que un programa haga más de una cosa a la vez. En Go, esto se logra utilizando goroutines y canales. Las goroutines son funciones que se ejecutan de manera concurrente con otras goroutines. Los canales permiten la comunicación segura entre goroutines.

### Goroutines

Una goroutine es una función que se ejecuta de manera concurrente con otras goroutines en el mismo espacio de direcciones. Las goroutines son muy ligeras y se crean utilizando la palabra clave `go`.

#### Ejemplo de Goroutines

```go
package main

import (
    "fmt"
    "time"
)

func imprimirNumeros() {
    for i := 1; i <= 5; i++ {
        fmt.Println(i)
        time.Sleep(100 * time.Millisecond)
    }
}

func main() {
    go imprimirNumeros() // Iniciar una goroutine

    // Continuar con la ejecución principal
    fmt.Println("Hola desde la goroutine principal")

    // Esperar para asegurarse de que la goroutine termine
    time.Sleep(600 * time.Millisecond)
}
```

**Explicación:**
- `go imprimirNumeros()`: Inicia la función `imprimirNumeros` como una goroutine.
- `time.Sleep(600 * time.Millisecond)`: Pausa la ejecución principal para permitir que la goroutine termine.

### Canales

Los canales son conduits que permiten la comunicación segura entre goroutines. Los canales en Go se crean utilizando la función `make` y se especifica el tipo de datos que pueden transportar.

#### Ejemplo de Canales

```go
package main

import (
    "fmt"
)

func productor(canal chan int) {
    for i := 1; i <= 5; i++ {
        canal <- i // Enviar datos al canal
    }
    close(canal) // Cerrar el canal al finalizar
}

func consumidor(canal chan int) {
    for valor := range canal { // Recibir datos del canal
        fmt.Println("Recibido:", valor)
    }
}

func main() {
    canal := make(chan int)

    go productor(canal)  // Iniciar goroutine productor
    consumidor(canal)    // Iniciar consumidor en la goroutine principal
}
```

**Explicación:**
- `canal := make(chan int)`: Crea un canal de tipo `int`.
- `canal <- i`: Envía el valor `i` al canal.
- `valor := range canal`: Recibe valores del canal hasta que esté cerrado.
- `close(canal)`: Cierra el canal para indicar que no se enviarán más valores.

### Select

La declaración `select` en Go permite esperar en múltiples operaciones de canal. `select` bloquea hasta que uno de sus casos pueda proceder, luego ejecuta ese caso.

#### Ejemplo de Select

```go
package main

import (
    "fmt"
    "time"
)

func enviar(canal chan string, mensaje string, retraso time.Duration) {
    time.Sleep(retraso)
    canal <- mensaje
}

func main() {
    canal1 := make(chan string)
    canal2 := make(chan string)

    go enviar(canal1, "mensaje de canal1", 2*time.Second)
    go enviar(canal2, "mensaje de canal2", 1*time.Second)

    for i := 0; i < 2; i++ {
        select {
        case mensaje1 := <-canal1:
            fmt.Println("Recibido:", mensaje1)
        case mensaje2 := <-canal2:
            fmt.Println("Recibido:", mensaje2)
        }
    }
}
```

**Explicación:**
- `select { ... }`: Permite esperar en múltiples operaciones de canal.
- `case mensaje1 := <-canal1`: Maneja el caso en que se recibe un mensaje de `canal1`.
- `case mensaje2 := <-canal2`: Maneja el caso en que se recibe un mensaje de `canal2`.

### Sincronización con `sync.WaitGroup`

El paquete `sync` proporciona un tipo `WaitGroup` que se utiliza para esperar a que un conjunto de goroutines termine.

#### Ejemplo de `sync.WaitGroup`

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

func trabajador(id int, wg *sync.WaitGroup) {
    defer wg.Done() // Indicar que la goroutine ha terminado
    fmt.Printf("Trabajador %d comenzando\n", id)
    time.Sleep(time.Second)
    fmt.Printf("Trabajador %d terminado\n", id)
}

func main() {
    var wg sync.WaitGroup

    for i := 1; i <= 3; i++ {
        wg.Add(1) // Incrementar el contador de WaitGroup
        go trabajador(i, &wg)
    }

    wg.Wait() // Esperar a que todas las goroutines terminen
    fmt.Println("Todos los trabajadores han terminado")
}
```

**Explicación:**
- `var wg sync.WaitGroup`: Declara un `WaitGroup`.
- `wg.Add(1)`: Incrementa el contador del `WaitGroup`.
- `defer wg.Done()`: Decrementa el contador del `WaitGroup` cuando la goroutine termina.
- `wg.Wait()`: Bloquea hasta que el contador del `WaitGroup` sea cero.

### Mutexes

El paquete `sync` también proporciona un tipo `Mutex` para exclusión mutua, que se utiliza para proteger el acceso a datos compartidos entre múltiples goroutines.

#### Ejemplo de Mutexes

```go
package main

import (
    "fmt"
    "sync"
)

var (
    contador int
    mutex    sync.Mutex
)

func incrementar(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        mutex.Lock()   // Bloquear el acceso al contador
        contador++
        mutex.Unlock() // Desbloquear el acceso al contador
    }
}

func main() {
    var wg sync.WaitGroup

    for i := 0; i < 10; i++ {
        wg.Add(1)
        go incrementar(&wg)
    }

    wg.Wait()
    fmt.Println("Contador final:", contador)
}
```

**Explicación:**
- `var mutex sync.Mutex`: Declara un `Mutex`.
- `mutex.Lock()`: Bloquea el acceso a una sección crítica.
- `mutex.Unlock()`: Desbloquea el acceso a una sección crítica.

### Resumen

En este tema, hemos cubierto los conceptos fundamentales de la concurrencia en Go, incluyendo el uso de goroutines, canales, `select`, `sync.WaitGroup` y mutexes. La concurrencia es una característica poderosa de Go que permite escribir programas eficientes y escalables. Comprender y aplicar estos conceptos te permitirá aprovechar al máximo las capacidades concurrentes de Go.