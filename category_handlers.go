package main

import (
        "fmt"
        "net/http"
        "encoding/json"
)

type Category struct {
        Description string `json:"description"`
}

func getCategoryHandler(w http.ResponseWriter, r *http.Request) {
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

func createCategoryHandler(w http.ResponseWriter, r *http.Request) {
        category := Category{}

        err := r.ParseForm()

        if err!= nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        category.Description = r.Form.Get("description")

        err = store.CreateCategory(&category)
        if err != nil {
                fmt.Println(err)
        }

        http.Redirect(w, r, "/a/", http.StatusFound)
}
