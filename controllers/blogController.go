package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/shandilyamanvesh/CrudAPIInGo/dal"
	"github.com/shandilyamanvesh/CrudAPIInGo/models"
	"net/http"
	"strconv"
)

//http handler for "/":get all blogs
func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	blogs := dal.GetAllBlogs()

	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//http handler for "/blogs":get all blogs
func BlogIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	blogs := dal.GetAllBlogs()

	if err := json.NewEncoder(w).Encode(blogs); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//http handler for "/blogs/{blogId}":get blog with a particular Id
func Show(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	blogId := vars["blogId"]

	id, err := strconv.Atoi(blogId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blog, err := dal.GetBlogById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(*blog); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//http handler for "/blogs":save a blog
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	blog := models.Blog{}
	err := decoder.Decode(&blog)
	if err != nil {
		// InvalidFormat error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blog = dal.CreateBlog(blog)

	if err := json.NewEncoder(w).Encode(blog); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

//http handler for "/blogs/{blogId}":update blog with at a particular Id
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	blogId := vars["blogId"]

	id, err := strconv.Atoi(blogId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	decoder := json.NewDecoder(r.Body)
	blog := models.Blog{}
	err = decoder.Decode(&blog)
	if err != nil {
		// InvalidFormat error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	blog.Id = id
	b, err := dal.UpdateBlogById(blog)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	if err := json.NewEncoder(w).Encode(*b); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

//http handler for "/blogs/{blogId}":delete blog with a particular Id
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	blogId := vars["blogId"]

	id, err := strconv.Atoi(blogId)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = dal.DeleteBlogById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
}
