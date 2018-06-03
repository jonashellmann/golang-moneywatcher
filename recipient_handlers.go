package main

import (
        "fmt"
        "net/http"
        "encoding/json"
)

type Recipient struct {
        Name string `json:"name"`
}

func getRecipientHandler(w http.ResponseWriter, r *http.Request) {
        recipients, err := store.GetRecipients()

        recipientListBytes, err := json.Marshal(recipients)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.Write(recipientListBytes)
}
