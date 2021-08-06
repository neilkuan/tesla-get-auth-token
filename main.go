package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	id := os.Getenv("PROJECT_ID")

	if id == "" {
		log.Fatalf("The Project ID is Null, Please setting PROJECT_ID environments variables.")
		os.Exit(1)
	}
	apiServer := gin.New()
	apiServer.GET("/hello", HelloWorldHTTP)

	apiServer.Run(":8080")

	HelloWorld(id)

}

// HelloWorld is to Say Hello World to User.
func HelloWorld(id string) {
	fmt.Println("Hello world!!!")
	fmt.Printf("Project id is %v \n", id)
}

// HelloWorldHTTP is to Say Hello World to User.
func HelloWorldHTTP(context *gin.Context) {
	context.JSON(http.StatusOK, "Hello World")
}
