package main

import(
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request){
if r.URL.Path != "/hello" {
	http.Error(w, "404 not found", http.StatusNotFound)
	
}
}

func main(){

	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8000\n")
	if err := http.ListenAndServe(":8000", nil); err !=nil (
		log.Fatal(err)
)
/}


