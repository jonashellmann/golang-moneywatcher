package main

import (
        "fmt"
        "net/http"
        "encoding/json"
)

type Recipient struct {
        Name   string `json:"name"`
	UserId int    `json:"userId"`
}

func getRecipientHandler(w http.ResponseWriter, r *http.Request) {
        userId, err := CheckCookie(r)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

	recipients, err := store.GetRecipients(userId)

        recipientListBytes, err := json.Marshal(recipients)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.Write(recipientListBytes)
}

func createRecipientHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	recipient := Recipient{}

        err = r.ParseForm()

        if err!= nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

	recipient.UserId = userId
        recipient.Name = r.Form.Get("name")

        err = store.CreateRecipient(&recipient)
        if err != nil {
                fmt.Println(err)
        }

        http.Redirect(w, r, "/a/", http.StatusFound)
}
