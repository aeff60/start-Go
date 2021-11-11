//https://docs.microsoft.com/en-us/learn/modules/serverless-go/3-web-api
package main

import (
	"fmt"
	"log"
	"net/http"
)
/*
http.HandleFunc("/", handleRoute)

func handleRequest(w: http:ResponseWriter, r: http.Request) {
    fmt.Fprintf(w, "My first REST API") 
}

http.ListenAndServe(":3000", nil)*/

func routeHandler(w http.ResponseWriter, r*http.Request){
	fmt.Println("route hit")
}

func main(){
	http.HandleFunc("/", routeHandler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}