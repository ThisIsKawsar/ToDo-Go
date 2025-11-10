package main

import (
    "github.com/gin-gonic/gin"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "github.com/gin-contrib/cors"
)

var db *gorm.DB
var err error

func main() {
    // MySQL connection
    dsn := "root:GW8ZmHPf@tcp(127.0.0.1:3306)/todo_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    db.AutoMigrate(&Todo{}) // must be inside main()

    r := gin.Default()
    r.Use(cors.Default()) // allow Vue frontend requests

    // Routes
    r.GET("/todos", GetTodos)
    r.POST("/todos", CreateTodo)
    r.PUT("/todos/:id", UpdateTodo)
    r.DELETE("/todos/:id", DeleteTodo)

    r.Run(":8081") // start server
}
