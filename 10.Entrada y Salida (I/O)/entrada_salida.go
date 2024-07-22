package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Manejo de Archivos
func manejoArchivos() {
	fmt.Println("=== Manejo de Archivos ===")
	archivo, err := os.Create("archivo.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo:", err)
		return
	}
	defer archivo.Close()

	_, err = archivo.WriteString("Hola, Go!\n")
	if err != nil {
		fmt.Println("Error al escribir en el archivo:", err)
		return
	}

	fmt.Println("Archivo creado y escrito con éxito")
}

// Lectura de Archivos
func lecturaArchivos() {
	fmt.Println("=== Lectura de Archivos ===")
	archivo, err := os.Open("archivo.txt")
	if err != nil {
		fmt.Println("Error al abrir el archivo:", err)
		return
	}
	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		fmt.Println("Línea:", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error al leer el archivo:", err)
	}
}

// Manipulación de Cadenas
func manipulacionCadenas() {
	fmt.Println("=== Manipulación de Cadenas ===")
	texto := "Hola, Go!"

	mayusculas := strings.ToUpper(texto)
	fmt.Println("Mayúsculas:", mayusculas)

	reemplazado := strings.ReplaceAll(texto, "Go", "Golang")
	fmt.Println("Reemplazado:", reemplazado)

	partes := strings.Split(texto, ", ")
	fmt.Println("Partes:", partes)

	contiene := strings.Contains(texto, "Go")
	fmt.Println("Contiene 'Go':", contiene)
}

// Trabajo con JSON
type Persona struct {
	Nombre string `json:"nombre"`
	Edad   int    `json:"edad"`
}

func trabajoConJSON() {
	fmt.Println("=== Trabajo con JSON ===")
	persona := Persona{Nombre: "Juan", Edad: 30}

	jsonData, err := json.Marshal(persona)
	if err != nil {
		fmt.Println("Error al codificar a JSON:", err)
		return
	}
	fmt.Println("JSON:", string(jsonData))

	err = os.WriteFile("persona.json", jsonData, 0644)
	if err != nil {
		fmt.Println("Error al escribir el archivo JSON:", err)
		return
	}

	data, err := os.ReadFile("persona.json")
	if err != nil {
		fmt.Println("Error al leer el archivo JSON:", err)
		return
	}

	var personaDecodificada Persona
	err = json.Unmarshal(data, &personaDecodificada)
	if err != nil {
		fmt.Println("Error al decodificar JSON:", err)
		return
	}
	fmt.Println("Persona decodificada:", personaDecodificada)
}

func main() {
	manejoArchivos()
	time.Sleep(100 * time.Millisecond) // Pausa para asegurar el orden en la salida
	lecturaArchivos()
	time.Sleep(100 * time.Millisecond) // Pausa para asegurar el orden en la salida
	manipulacionCadenas()
	time.Sleep(100 * time.Millisecond) // Pausa para asegurar el orden en la salida
	trabajoConJSON()
}
