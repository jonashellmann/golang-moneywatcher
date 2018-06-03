package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type Region struct {
	Description string `json:"description"`
}

func getRegionHandler(w http.ResponseWriter, r *http.Request) {
	regions, err := store.GetRegions()

	regionListBytes, err := json.Marshal(regions)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(regionListBytes)
}
