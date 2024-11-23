package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type todo struct {
	ID        string `json:"id"` // ID should always be in capital letters
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Learn Go", Completed: false},
	{ID: "2", Item: "Learn Gin", Completed: false},
	{ID: "3", Item: "Learn React", Completed: false},
	{ID: "4", Item: "Learn Angular", Completed: false},
	{ID: "5", Item: "Learn Vue", Completed: false},
	{ID: "6", Item: "Learn Node", Completed: false},
	{ID: "7", Item: "Learn Express", Completed: false},
}

func getTodos(context *gin.Context) {
	fmt.Printf("type of context %T\n", context)
	context.IndentedJSON(http.StatusOK, todos) // convert to json
}

func addTodo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil { // BindJSON is a method that binds the request body to a struct
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func getTodosById(context *gin.Context) {
	id := context.Param("id")
	for _, todo := range todos {
		if todo.ID == id {
			context.IndentedJSON(http.StatusOK, todo)
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func toggleTodos(context *gin.Context) {
	id := context.Param("id")
	for index, todo := range todos {
		if todo.ID == id {
			todos[index].Completed = !todos[index].Completed
			context.IndentedJSON(http.StatusOK, todos[index])
			return
		}
	}
	context.IndentedJSON(http.StatusNotFound, gin.H{"message": "todo not found"})
}

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func main() {
	go recordMetrics() // Start recording metrics in a separate goroutine

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil) // Start Prometheus endpoint

	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.GET("/todos/:id", getTodosById)
	router.PATCH("/todos/:id", toggleTodos)
	router.Run("localhost:9090") // Start Gin server
}
