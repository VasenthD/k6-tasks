package main

import (
    "context"
    "fmt"
    "log"
    "net/http"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

    "github.com/gin-gonic/gin"
)

// Model representing the document structure
type K6Document struct {
    Value int `bson:"value"`
}

func main() {
    r := gin.Default()

    // Define a route for inserting a document
    r.POST("/api/insert", func(c *gin.Context) {
        // Set your MongoDB connection string
        connectionString := "mongodb+srv://vasenth01:ved@college.yx7msef.mongodb.net/"

        // Initialize MongoDB client options
        clientOptions := options.Client().ApplyURI(connectionString)

        // Connect to MongoDB
        client, err := mongo.Connect(context.Background(), clientOptions)
        if err != nil {
            log.Fatal(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to connect to the database"})
            return
        }

        // Ensure proper disconnection when the request completes
        defer client.Disconnect(context.Background())

        // Get a handle to the "k6" collection
        collection := client.Database("college").Collection("k6")

        // Create a document to insert
        document := K6Document{Value: 1}

        // Insert the document into the collection
        _, err = collection.InsertOne(context.Background(), document)
        if err != nil {
            log.Fatal(err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document into the database"})
            return
        }

        c.JSON(http.StatusOK, gin.H{"message": "Document inserted successfully"})
    })

    fmt.Println("Server running on port 2000")
    log.Fatal(r.Run(":2000"))
}
