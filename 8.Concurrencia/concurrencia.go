package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    // Goroutines
    go saludar("Mundo")
    time.Sleep(1 * time.Second) // Esperar a que la goroutine termine

    // Canales
    mensajes := make(chan string)
    go enviarMensaje(mensajes)
    mensaje := <-mensajes
    fmt.Println("Mensaje recibido:", mensaje)

    // Select
    c1 := make(chan string)
    c2 := make(chan string)
    go func() {
        time.Sleep(1 * time.Second)
        c1 <- "uno"
    }()
    go func() {
        time.Sleep(2 * time.Second)
        c2 <- "dos"
    }()

    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("Recibido", msg1)
        case msg2 := <-c2:
            fmt.Println("Recibido", msg2)
        }
    }

    // Mutex
    var contador int
    var mutex sync.Mutex
    var wg sync.WaitGroup
    wg.Add(2)

    go incrementar(&contador, &mutex, &wg)
    go incrementar(&contador, &mutex, &wg)

    wg.Wait()
    fmt.Println("Contador final:", contador)
}

func saludar(nombre string) {
    fmt.Println("Hola", nombre)
}

func enviarMensaje(mensajes chan string) {
    mensajes <- "Hola desde la goroutine"
}

func incrementar(contador *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 1000; i++ {
        mutex.Lock()
        *contador++
        mutex.Unlock()
    }
}
