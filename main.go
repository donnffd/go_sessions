package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("this-has-to-be-a-secret-key"))

func secret(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")

	if auth , ok := session.Values["authentication"]; !ok || !auth.(bool) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	fmt.Fprintln(w, "The cake is a lie!")
}

func login(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")

	/* @todo Authentication goes here */

	session.Values["authentication"] = true


	session.Save(r, w)

}

func logout(w http.ResponseWriter, r *http.Request){
	session, _ := store.Get(r, "cookie-name")

	//@todo Revoke users authenication

	session.Values["authentication"] = false


	session.Save(r, w)
}

func main(){
	http.HandleFunc("/secret", secret)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)

	http.ListenAndServe(":8080",nil)

}