package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// Get all todos
func GetTodos(c *gin.Context) {
    var todos []Todo
    db.Find(&todos)
    c.JSON(http.StatusOK, todos)
}

// Create a new todo
func CreateTodo(c *gin.Context) {
    var todo Todo
    if err := c.ShouldBindJSON(&todo); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&todo)
    c.JSON(http.StatusOK, todo)
}

// Update todo by ID
func UpdateTodo(c *gin.Context) {
    id := c.Param("id")
    var todo Todo
    if err := db.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }

    var input Todo
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    todo.Title = input.Title
    todo.Completed = input.Completed
    db.Save(&todo)

    c.JSON(http.StatusOK, todo)
}

// Delete todo by ID
func DeleteTodo(c *gin.Context) {
    id := c.Param("id")
    var todo Todo
    if err := db.First(&todo, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
        return
    }
    db.Delete(&todo)
    c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
