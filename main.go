package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	id := os.Getenv("PROJECT_ID")

	if id == "" {
		log.Fatalf("The Project ID is Null, Please setting PROJECT_ID environments variables.")
		os.Exit(1)
	}
	apiServer := gin.New()
	apiServer.GET("/hello", HelloWorldHTTP)
	apiServer.GET("/getallvehicles", GetAllVehicles)

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

// GetAllVehicles is to Say Hello World to User.
func GetAllVehicles(context *gin.Context) {
	url := "https://owner-api.teslamotors.com/api/1/vehicles"

	// Create a Bearer string by appending string access token
	var bearer = "Bearer " + os.Getenv("ACCESS_TOKEN")

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	// add authorization header to the req
	req.Header.Add("Authorization", bearer)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	log.Println(string([]byte(body)))
	context.JSON(http.StatusOK, string([]byte(body)))
}
