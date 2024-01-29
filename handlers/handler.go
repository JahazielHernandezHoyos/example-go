package handlers

import (
    "github.com/gin-gonic/gin"
    "nombre_del_proyecto/models"
    "net/http"
)

var items []models.Item

// Obtener todos los elementos
func GetItems(c *gin.Context) {
    c.JSON(http.StatusOK, items)
}

// Obtener un elemento por ID
func GetItem(c *gin.Context) {
    // Lógica para obtener un elemento por ID
}

// Crear un nuevo elemento
func CreateItem(c *gin.Context) {
    // Lógica para crear un nuevo elemento
}

// Actualizar un elemento por ID
func UpdateItem(c *gin.Context) {
    // Lógica para actualizar un elemento por ID
}

// Eliminar un elemento por ID
func DeleteItem(c *gin.Context) {
    // Lógica para eliminar un elemento por ID
}
