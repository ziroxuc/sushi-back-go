package rutas

import (
	"github.com/gorilla/mux"
	"net/http"
	"../actions"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Name(route.Name).
			Path(route.Pattern).
			Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index", "GET", "/", actions.Index},
	//Route{
	//	"Index", "GET", "/persona/{id}", actions.GetPersona},
	Route{
		"getPersona", "GET", "/personas", actions.PersonaList},
	Route{
		"addPersona", "POST", "/persona", actions.AddPersona},
	Route{
		"GetPersonaById", "GET", "/persona/{id}", actions.GetPersonaById},

	Route{
		"GetPersonaid2", "GET", "/personas/{id2}", actions.GetPersonaId2},



}
