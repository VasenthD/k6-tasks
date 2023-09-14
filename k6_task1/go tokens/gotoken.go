package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	collection *mongo.Collection
)

func init() {

	clientOptions := options.Client().ApplyURI("mongodb+srv://vasenth01:ved@college.yx7msef.mongodb.net/")
	client, _ = mongo.Connect(context.TODO(), clientOptions)

	// Check the connection
	err := client.Ping(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	
	database := client.Database("college")        
	collection = database.Collection("gotokens") 
}

func Token(token string, c *gin.Context) {
	_, err := collection.InsertOne(context.TODO(), map[string]interface{}{"token": token})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Token stored successfully"})
}


func main() {
	router := gin.Default()
fmt.Println("routes starting...")
	

	router.POST("/tokens", func(c *gin.Context) {
		
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Token not found in the header"})
			return
		}
		go Token(token,c)
		
	})

	



	router.Run(":8080")
}