package controllers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("index")
}

func BlogIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("blogindex")
}

func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("show")
}

func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("save")
}
