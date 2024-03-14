// Lista de Tareas (Todo List):
// Construye una aplicación de lista de tareas que permita a los usuarios agregar, completar y eliminar tareas.

// 1. Buscar framework para hacer la aplicacion
// 2. Empezar a crear la logica
// 2.1 Hago uso del metodo get para visualizar todas las tareas
// 2.2 Hago uso del todo post para crear tareas en un nuevo slice del tipo tarea
// 2.3 Hago uso del metodo get para tener la lista completa pero uso un for para solo mostrar los id iguales
// 2.4 Hago uso del metodo path para cambiar temporalmente el completed
// 2.5 Hago uso del metodo delete para eliminar el recurso(id) que me pida el usuario

// package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type tarea struct {
// 	ID        string `json:"id"`
// 	Task      string `json:"task"`
// 	Completed bool   `json:"completed"`
// }

// var tareas = []tarea{
// 	{ID: "1", Task: "Estudiar", Completed: true},
// 	{ID: "2", Task: "Pasear", Completed: false},
// 	{ID: "3", Task: "Jugar", Completed: true},
// }

// func getTareas(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, tareas)
// }

// func postTareas(c *gin.Context) {
// 	var newTarea tarea

// 	err := c.BindJSON(&newTarea)
// 	if err != nil {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//         return
// 	}

// 	tareas = append(tareas, newTarea)
// 	c.IndentedJSON(http.StatusCreated, newTarea)

// }

// func getTareasByID(c *gin.Context) {
// 	id := c.Param("id")

// 	for _, a := range tareas {
// 		if a.ID == id {
// 			c.IndentedJSON(http.StatusOK, a)
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
// }

// func patchTareasCompleted(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, task := range tareas {
// 		if task.ID == id {
// 			tareas[i].Completed = !tareas[i].Completed
// 			c.IndentedJSON(http.StatusOK, tareas[i])
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
// }

// func deleteTareas(c *gin.Context) {
// 	id := c.Param("id")

// 	for i, task := range tareas {
// 		if task.ID == id {
// 			tareas = append(tareas[:i], tareas[i+1:]...)
// 			c.IndentedJSON(http.StatusOK, gin.H{"message": "Tarea eliminada correctamente"})
// 			return
// 		}
// 	}

// 	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
// }

// func main() {
// 	router := gin.Default()
// 	router.GET("/tareas", getTareas)
// 	router.GET("/tareas/:id", getTareasByID)
// 	router.POST("/tareas", postTareas)
// 	router.PATCH("/tareas/:id/completed", patchTareasCompleted)
// 	router.DELETE("/tareas/:id", deleteTareas)

// 	router.Run("localhost:8080")
// }
