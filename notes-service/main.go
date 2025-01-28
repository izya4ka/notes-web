package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/notes", func(c *gin.Context) {
        c.String(http.StatusOK, "notes")
    })
    
    router.POST("/notes", func (c *gin.Context) {
        c.String(http.StatusOK, "notes post")
    })

    router.GET("/notes/:id", func (c *gin.Context) {
        c.String(http.StatusOK, "notes " + c.Param("id") + " get")
    })

    router.PATCH("/notes/:id", func (c *gin.Context) {
        c.String(http.StatusOK, "notes" + c.Param("id") + " patch")
    })

    router.DELETE("/notes/:id", func (c *gin.Context) {
        c.String(http.StatusOK, "notes" + c.Param("id") + " delete")
    })
}