package routes

import (
    "github.com/gin-gonic/gin"
    "app/handlers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    // Rutas para el CRUD
    r.GET("/items", handlers.GetItems)
    r.GET("/items/:id", handlers.GetItem)
    r.POST("/items", handlers.CreateItem)
    r.PUT("/items/:id", handlers.UpdateItem)
    r.DELETE("/items/:id", handlers.DeleteItem)

    return r
}
