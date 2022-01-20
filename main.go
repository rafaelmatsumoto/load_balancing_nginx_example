package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main()  {
	log.Print("Server loaded")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "APPID = %s\n", os.Getenv("APPID"))
}
