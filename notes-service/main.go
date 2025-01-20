package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": os.Getenv("DB_URL")})
    })

    router.GET("/users/:id", getUserByID)

    router.Run(":8080")
}

func getUserByID(c *gin.Context) {
    id := c.Param("id")
    c.JSON(http.StatusOK, gin.H{"userID": id})
}
