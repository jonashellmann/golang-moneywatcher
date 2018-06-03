package main

import (
        "fmt"
        "net/http"
        "encoding/json"
)

type Expense struct {
        Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Date        string  `json:"date"`
	CategoryId  int     `json:"categoryId"`
	RegionId    int     `json:"regionId"`
	RecipientId int     `json:"recipientId"`
}

func getExpenseHandler(w http.ResponseWriter, r *http.Request) {
        expenses, err := store.GetExpenses()

        expenseListBytes, err := json.Marshal(expenses)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.Write(expenseListBytes)
}
