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
