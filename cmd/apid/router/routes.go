package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mateusmatinato/goexpert-cep2temp/internal/cep2temp"
)

func SetupRouter(handler cep2temp.Handler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/cep2temp/{cep}", handler.GetTemperatureByCEP).Methods(http.MethodGet)
	return r
}
