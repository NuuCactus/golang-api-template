package router

import (
	_ "github.com/nuucactus/golang-api-template/endpoints/metrics"
	"github.com/nuucactus/golang-api-template/endpoints/health"

	"github.com/charmixer/oas/api"
	"github.com/charmixer/oas/exporter"

	"github.com/gorilla/mux"
)

var oas api.Api

func init() {
	oas = api.Api{}

	oas.Title = "Golang api template"
	oas.Description = `Gives a simple blueprint for creating new api's`
	oas.Version = "0.0.0"
	oas.NewPath("GET", "/health", health.GetHealth, health.GetHealthSpec())

	exporter.ToOasModel(oas)
}


func NewRouter() (r *mux.Router) {
	r = mux.NewRouter()

	for _,p := range oas.GetPaths() {
		r.HandleFunc(p.Url, p.Handler).Methods(p.Method)
	}

	return r
}
