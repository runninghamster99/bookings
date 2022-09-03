package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

/*
function that takes a handler and returns hander
*/

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		//move on to the next, another middleware or where we return our mux
		next.ServeHTTP(w, r)
	})
}

// all middleware returns http.Handler, they must return this
// adds CSRF protection for all POST request
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)

	//uses cookie to make sure token we generate is available on per page basis
	//creating new cookie with http.Cookie, and just enough information for noSurf
	csrfHandler.SetBaseCookie(http.Cookie{

		HttpOnly: true,
		Path:     "/",                  //to apply to entire site, using path "/"
		Secure:   app.InProduction,     //we are using http and not https
		SameSite: http.SameSiteLaxMode, //built in standard SameSiteLaxMode

	})
	return csrfHandler
}

//provides middleware which autmatically loads and saves session data for the current request,
// and communicates the session token to and from the client in a cookie
// web servers by their very nature are not state aware, i need to add middleware to tell it
// to remember state using sessions, new middleware for it

//SessionLoad loads and saves the session on every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
