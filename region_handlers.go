package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Region struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	UserId      int    `json:"userId"`
}

func getRegionHandler(w http.ResponseWriter, r *http.Request) {
        userId, err := CheckCookie(r)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

	regions, err := store.GetRegions(userId)

	regionListBytes, err := json.Marshal(regions)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(regionListBytes)
}

func createRegionHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	region := Region{}

	err = r.ParseForm()

	if err!= nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	region.UserId = userId
	region.Description = r.Form.Get("description")

	err = store.CreateRegion(&region)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/a/", http.StatusFound)
}
