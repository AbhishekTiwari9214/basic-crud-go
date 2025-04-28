package main

import (
	"basic-crud/utils"
	"fmt"
	"log"
	"net/http"
)


func main(){
	utils.LoadEnv()
	port := utils.GetEnv("PORT") ;
	fmt.Println("running the server at the port:" + port)
	if err := http.ListenAndServe(":"+port ,nil); err!=nil {
		log.Fatal("error while starting the server:",err)
	}
}