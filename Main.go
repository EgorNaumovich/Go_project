package main

import (
	"awesomeProject2/Go_project/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	l := log.New(os.Stdout, "test", log.LstdFlags)
	productHandler := handlers.ProductsLog(l)
	s := http.NewServeMux()
	s.Handle("/", productHandler)
	http.ListenAndServe(":8080", s)
}
