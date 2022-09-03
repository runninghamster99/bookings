package models

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	//cross site request forgery token. hidden random numbers that change everytime one goes
	//the page. go lang no surf package can do this 
	//generates CSRF for us and prevent us github.com/justinas/nosurf
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}
