package main

import (
	"github.com/gorilla/securecookie"
	"net/http"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

func loginHandler(response http.ResponseWriter, request *http.Request) {
	username := request.Form.Get("username")
	password := request.Form.Get("password")

	redirectTarget := "/a/login.html?login=false"
	if username != "" && password != "" {
		// TODO: Check credentials
		loginValid := store.CheckCredentials(username, password)
		if loginValid {
			setSession(username, response)
			redirectTarget = "/a/"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

func logoutHandler(response http.ResponseWriter, request *http.Request) {
	clearSession(response)
	http.Redirect(response, request, "/a/login.html?loggedOut=true", 302)
}

func setSession(userName string, response http.ResponseWriter) {
	value := map[string]string {
		"name": userName,
	}

	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
	}
}

func clearSession(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

func CheckCookie(request *http.Request) (int, error) {
	cookie, err := request.Cookie("session")
	if err == nil {
		cookieValue := make(map[string]string)
		err = cookieHandler.Decode("session", cookie.Value, &cookieValue)
		if err == nil {
			username := cookieValue["name"]
			userId, err := store.GetUserId(username)
			if err == nil {
				return userId, err
			}
		}
	}
	return 0, err
}
