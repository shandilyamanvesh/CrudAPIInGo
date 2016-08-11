package controllers

import (
	"fmt"
	"net/http"
)

//http handler for "/":get all blogs
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("index")
}

//http handler for "/blogs":get all blogs
func BlogIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("blogindex")
}

//http handler for "/blogs/{blogId}":get blog with a particular Id
func Show(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("show")
}

//http handler for "/blogs":save a blog
func Save(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("save")
}
