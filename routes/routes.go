package routes

import (
	"net/http"

	"Tracky/controller"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/trackable", controller.LogTrackable).Methods(http.MethodPost)

	return router
}
