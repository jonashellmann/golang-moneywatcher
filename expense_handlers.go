package main

import (
        "fmt"
        "net/http"
        "encoding/json"
	"database/sql"
	"github.com/go-sql-driver/mysql"
)

type Expense struct {
        Description sql.NullString    `json:"description"`
	Amount      float64   `json:"amount"`
	Date        mysql.NullTime `json:"date"`
	CategoryId  sql.NullInt64       `json:"categoryId"`
	RegionId    sql.NullInt64       `json:"regionId"`
	RecipientId sql.NullInt64       `json:"recipientId"`
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
