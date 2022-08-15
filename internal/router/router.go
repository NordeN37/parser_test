package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"parser_test/internal/router/parser_controller"
)

func InitRouter(router *mux.Router) {
	sr := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	sr.HandleFunc("/parsers", wrapJSONHandler(parser_controller.Parser)).Methods(http.MethodPost)
}
