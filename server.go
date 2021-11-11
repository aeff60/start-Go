package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
	"os"
 )

 func helloRequest(w http.ResponseWriter, r*http.Request) {
    //fmt.Fprintf(w, "My first REST AP") 
	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		w.Write([]byte("hello world"))
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		w.Write(body)
	}
}

func main(){
	//http.HandleFunc("/", handleRequest)
	//http.ListenAndServe(":3000", nil)
	customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
   		customHandlerPort = "8080"
	} // 1)
	mux := http.NewServeMux() // 2)
	mux.HandleFunc("/hello", helloRequest) // 3)
	fmt.Println("Go server Listening on: ", customHandlerPort)
	log.Fatal(http.ListenAndServe(":"+customHandlerPort, mux)) // 4)
}
