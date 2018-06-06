package main

import (
        "fmt"
        "net/http"
        "encoding/json"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"strconv"
	"time"
)

type Expense struct {
        Description sql.NullString `json:"description"`
	Amount      float64        `json:"amount"`
	Date        mysql.NullTime `json:"date"`
	CategoryId  sql.NullInt64  `json:"categoryId"`
	RegionId    sql.NullInt64  `json:"regionId"`
	RecipientId sql.NullInt64  `json:"recipientId"`
}

func getExpenseHandler(w http.ResponseWriter, r *http.Request) {
        userId, err := CheckCookie(r)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

	expenses, err := store.GetExpenses(userId)

        expenseListBytes, err := json.Marshal(expenses)

        if err != nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

        w.Write(expenseListBytes)
}

func createExpenseHandler(w http.ResponseWriter, r *http.Request) {
	expense := Expense{}

        err := r.ParseForm()

        if err!= nil {
                fmt.Println(fmt.Errorf("Error: %v", err))
                w.WriteHeader(http.StatusInternalServerError)
                return
        }

	expense.Description = sql.NullString{String: r.Form.Get("description"), Valid: true}
	expense.Amount, _ = strconv.ParseFloat(r.Form.Get("amount"), 64)

	date, err := time.Parse(time.RFC3339, r.Form.Get("date"))
	if err == nil {
		expense.Date = mysql.NullTime{Time: date, Valid: true}
	} else {
		expense.Date = mysql.NullTime{Valid: false}
	}

	regionId, err := strconv.ParseInt(r.Form.Get("region"), 10, 64)
	if err == nil {
		expense.RegionId = sql.NullInt64{Int64: regionId, Valid: true}
	} else {
		expense.RegionId = sql.NullInt64{Valid: false}
	}

	recipientId, err := strconv.ParseInt(r.Form.Get("recipient"), 10, 64)
        if err == nil {
		expense.RecipientId = sql.NullInt64{Int64: recipientId, Valid: true}
	} else {
		expense.RecipientId = sql.NullInt64{Valid: false}
	}

	categoryId, err := strconv.ParseInt(r.Form.Get("category"), 10, 64)
        if err == nil {
		expense.CategoryId = sql.NullInt64{Int64: categoryId, Valid: true}
	} else {
		expense.CategoryId = sql.NullInt64{Valid: false}
	}

        err = store.CreateExpense(&expense)
        if err != nil {
                fmt.Println(err)
        }

        http.Redirect(w, r, "/a/", http.StatusFound)
}