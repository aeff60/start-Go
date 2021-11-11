package main

import (
	"fmt"
	"net/http"
 )



 func handleRequest(w http.ResponseWriter, r*http.Request) {
    fmt.Fprintf(w, "Hello") 
}

func main(){
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":3000", nil)
}
