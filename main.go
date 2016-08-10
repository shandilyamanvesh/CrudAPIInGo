package main

import (
	"fmt"
	"github.com/shandilyamanvesh/CrudAPIInGo/router"
	"log"
	"net/http"
)

func main() {
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
