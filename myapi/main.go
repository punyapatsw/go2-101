package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status string `json:"status"`
}

// var todos = map[int]*Todo{
// 	1: &Todo{ID: 1, Title: "pay phone bill", Status: "active"},
// }
var todos = map[int]*Todo{}

func helloHandler(c *gin.Context) {
	c.String(http.StatusOK, "hi.")
	// c.JSON(http.StatusOK, gin.H{
	// 	"message": "hello",
	// })
}

func getTodosHandler(c *gin.Context) {
	// todo := Todo{ID: 1, Title: "pun", Status: "0"}
	c.JSON(http.StatusOK, todos)
}

func getTodosByIdHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	t, ok := todos[id]
	if !ok {
		c.JSON(http.StatusOK, gin.H{})
		return
	}
	c.JSON(http.StatusOK, t)
}

func createTodosHandler(c *gin.Context) {
	t := Todo{}
	if err := c.ShouldBindJSON(&t); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idx := len(todos)
	idx++
	t.ID = idx
	todos[idx] = &t

	c.JSON(http.StatusCreated, "created todo.")
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(authMidlleware)
	r.GET("/hello", helloHandler)
	r.GET("/todos", getTodosHandler)
	r.GET("/todos/:id", getTodosByIdHandler)
	r.POST("/todos", createTodosHandler)
	return r
}

func authMidlleware(c *gin.Context) {
	log.Println("start middleware")
	authKey := c.GetHeader("Authorization")
	if authKey != "Bearer token123" {
		c.JSON(http.StatusBadRequest, "wrong token")
		c.Abort()
		return
	}

	c.Next()
	log.Println("end middleware")
}

func main() {
	r := setupRouter()
	r.Run(":1234") // listen and serve on 127.0.0.0:1234
}
