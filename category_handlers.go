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
        categorys, err := store.GetCategorys()

        categoryListBytes, err := json.Marshal(categorys)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.Write(categoryListBytes)
}
