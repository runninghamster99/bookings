package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/runninghamster99/bookings/internal/config"
	"github.com/runninghamster99/bookings/internal/handlers"
)


/*
take argument AppConfig and return http.Handler
A new http handler called mux or multiplexer using third party package and not built to function
pat and chi are two good routing, we are using chi
go get github.com/bmizerany/pat to get it, in go mod package will apear
return same handler, mux
middleware process requests as they come in and make decision about what to do with them. 
See whether user have ability to view certain pages. 
go get github.com/go-chi/chi

*/
func Routes(app *config.AppConfig) http.Handler {
	//a multipliexer which is an http.Handler
	mux := chi.NewRouter()

	//Gracefull absorb panics and prints the stack trace. 
	//Recover from panic and create a report about what happened
	mux.Use(middleware.Recoverer)

	//using middleware function write to console from middleware.go
	//mux.Use(WriteToConsole)
	mux.Use(NoSurf)

	//using middleware we made for sessions
	mux.Use(SessionLoad)


	//set up routes. mux has some methods built in. We want a get request
	//takes two arguments, pattern or string we are matching and http.HandlerFunc
	//handlers.Repo.Home, this will route to our home function in handler package
	//and so will work about
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	mux.Get("/generals-quarters", handlers.Repo.Generals)
	mux.Get("/majors-suite", handlers.Repo.Majors)
	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/search-availability-json", handlers.Repo.AvailabilityJSON)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
    //how to handle static file, javascript and css and images from local files
	//fileServer is a place to go and get static files from 
	fileServer := http.FileServer(http.Dir("./static/")) //dir give us directory from ./static/
	//using fileServer. we are going to handle, look for path name in static/ anything in the directory. StripPrefix takes URL which go gets and modify it for web server/client request into something it knows how to handle
	//string static and get fileServer
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
