package router

import (
	_ "github.com/nuucactus/golang-api-template/endpoints/metrics"
	"github.com/nuucactus/golang-api-template/endpoints/health"

	"github.com/charmixer/oas/api"

	"github.com/gorilla/mux"
)

func NewOas() (oas api.Api){
	oas = api.Api{}

	oas.Title = "Golang api template"
	oas.Description = `Gives a simple blueprint for creating new api's`
	oas.Version = "0.0.0"

	oas.NewPath("GET", "/health", health.GetHealth, health.GetHealthSpec())
	oas.NewPath("POST", "/health", health.PostHealth, health.PostHealthSpec())

	return oas
}


func NewRouter(oas api.Api) (r *mux.Router) {
	r = mux.NewRouter()

	for _,p := range oas.GetPaths() {
		r.HandleFunc(p.Url, p.Handler).Methods(p.Method)
	}

	return r
}
