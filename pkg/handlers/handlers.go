package handlers

import (
	"net/http"

	"github.com/runninghamster99/bookings/pkg/config"
	"github.com/runninghamster99/bookings/pkg/models"
	"github.com/runninghamster99/bookings/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	//playing with session, taking out ip of guy visit our address
	//remote IP address, pull it right of the request
	remoteIP := r.RemoteAddr //part into standard library part of http package
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	//pull value out of session, ip remote ip in home handler funciton above
	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")

	//m.App.Session. do more things

	stringMap["remote_ip"] = remoteIP

	//putting  a session here
	//m.App.Session

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}
