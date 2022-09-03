package main

import (
	"fmt"
	"net/http"
	//"pkg/handlers"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/runninghamster99/bookings/pkg/config"
	"github.com/runninghamster99/bookings/pkg/handlers"
	"github.com/runninghamster99/bookings/pkg/render"

	"log"
)

const portNumber = ":8080"

//cutted so it's availabe in middleware
var app config.AppConfig

//so it can be used in middleware
var session *scs.SessionManager

// main is the main function
func main() {

	//setting InProdution, change it to true when in production, there is still a better way
	//only need to put this true when in production
	app.InProduction = false

	//creating session with scs, this is defaults
	//this session is stored in cookies

	session = scs.New() //first sesion := scs.New() when writng session, changed for putting session variable in middleware
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true // cookie presist after closing the browser
	session.Cookie.SameSite = http.SameSiteLaxMode // how strict you want to be about what site you want this cookie applies to
	//session.Cookie.Secure = false converted to app.InProduction for managing only one value
	session.Cookie.Secure = app.InProduction //https and cookies are encrypted. false because localhost is being used
    //second time we had to use false, in middleware we set another cookie for no surf to be false
	// two things to be remembered
	//set this variable wherever it needs to be set only once, put it in AppConfig

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf("starting application on port %s", portNumber)
	//fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	//http.HandleFunc("/",handlers.Repo.Home)
	//http.HandleFunc("/about",handlers.Repo.Home)
	//We no longer need above built in routing, we will use make a routes.go file and use
	//third party package there

	//a new variable, which is serving, it will point to http server
	//Address, which is port number, and our handlers.
	//call routes and pass &app
	srv := &http.Server{
		Addr:    portNumber,
		Handler: Routes(&app),
	}

	//starting the server

	err = srv.ListenAndServe()
	//if an error
	if err != nil {
		log.Fatal(err)
	}
}





//used previously
// func main(){
	
// 	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
// 		_, _ = fmt.Fprintf(w, "Hello World!")
// 	})

// 	_ = http.ListenAndServe(":6677",nil)

// }

// import (
// 	"fmt"
// 	"net/http"
// )