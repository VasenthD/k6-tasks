package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Model representing the document structure
type K6Document struct {
	Value string `bson:"value"`
}

func main() {
	r := gin.Default()

	// Define a route for inserting a document
	r.POST("/ping", func(c *gin.Context) {
        var doc K6Document
        if err := c.BindJSON(&doc); err!= nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return 
        }

		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	fmt.Println("Server running on port 2000")
	log.Fatal(r.Run(":2000"))
}
