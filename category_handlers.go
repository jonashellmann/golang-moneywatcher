package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
)

type Category struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
}

func getCategorysHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	categorys, err := store.GetCategorys(userId)
	categoryListBytes, err := json.Marshal(categorys)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(categoryListBytes)
}

func getCategoryHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	vars := mux.Vars(r)
	categoryId, err := strconv.ParseInt(vars["categoryId"], 10, 64)
	
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	category, err := store.GetCategory(userId, categoryId)
	
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	categoryBytes, err := json.Marshal(category)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(categoryBytes)
}

func createCategoryHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	category := Category{}

	err = r.ParseForm()

	if err!= nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	category.UserId = userId
	category.Description = r.Form.Get("description")

	err = store.CreateCategory(&category)
	if err != nil {
			fmt.Println(err)
	}

	http.Redirect(w, r, "/a/", http.StatusFound)
}
