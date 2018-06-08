package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"strconv"
	"time"
	"strings"
	"github.com/gorilla/mux"
)

type Expense struct {
        Description sql.NullString `json:"description"`
	Amount      float64        `json:"amount"`
	Date        mysql.NullTime `json:"date"`
	Category    Category       `json:"category"`
	Region      Region         `json:"region"`
	Recipient   Recipient      `json:"recipient"`
	UserId      int            `json:"userId"`
}

func getExpensesHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
	}

	expenses, err := store.GetExpenses(userId)

if err != nil {
                        fmt.Println(fmt.Errorf("Error: %v", err))
                        w.WriteHeader(http.StatusInternalServerError)
                        return
        }	

	expenseListBytes, err := json.Marshal(expenses)

	if err != nil {
			fmt.Println(fmt.Errorf("Error: %v", err))
			w.WriteHeader(http.StatusInternalServerError)
			return
	}

	w.Write(expenseListBytes)
}

func getExpenseHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	vars := mux.Vars(r)
	expenseId, err := strconv.Atoi(vars["expenseId"])
	
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	expense, err := store.GetExpense(userId, expenseId)
	
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	expenseBytes, err := json.Marshal(expense)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(expenseBytes)
}

func createExpenseHandler(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookie(r)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expense := Expense{}

	err = r.ParseForm()

	if err!= nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	expense.UserId = userId
	expense.Description = sql.NullString{String: r.Form.Get("description"), Valid: true}
	expense.Amount, _ = strconv.ParseFloat(r.Form.Get("amount"), 64)

	dateString := r.Form.Get("date")
	dateString = strings.Replace(dateString, "-", "", -1)
	date, err := time.Parse("20060102", dateString)
	if err == nil {
		expense.Date = mysql.NullTime{Time: date, Valid: true}
	} else {
		fmt.Println(fmt.Errorf("Error: %v", err))
		expense.Date = mysql.NullTime{Valid: false}
	}

	regionId, err := strconv.Atoi(r.Form.Get("region"))
	if err == nil {
		expense.Region.Id = regionId
	} else {
		expense.Region.Id = 0
	}

	recipientId, err := strconv.Atoi(r.Form.Get("recipient"))
        if err == nil {
		expense.Recipient.Id = recipientId
	} else {
		expense.Recipient.Id = 0
	}

	categoryId, err := strconv.Atoi(r.Form.Get("category"))
        if err == nil {
		expense.Category.Id = categoryId
	} else {
		expense.Category.Id = 0
	}

	err = store.CreateExpense(&expense)
	if err != nil {
			fmt.Println(err)
	}

	http.Redirect(w, r, "/a/", http.StatusFound)
}
