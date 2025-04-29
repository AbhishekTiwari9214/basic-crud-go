package main

import (
	"basic-crud/routes"
	"basic-crud/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)


func main(){
	utils.LoadEnv()
	port := utils.GetEnv("PORT") ;
	client, err:= utils.GetMongoDbClient()
	if err!=nil {
		log.Fatal("db conection not connected")
	}
	defer utils.DisconnectDB(client)

	router := mux.NewRouter()
	routes.RegisterRoutes(router)

	fmt.Println("running the server at the port:" + port)
	if err := http.ListenAndServe(":"+port ,router); err!=nil {
		log.Fatal("error while starting the server:",err)
	}
}