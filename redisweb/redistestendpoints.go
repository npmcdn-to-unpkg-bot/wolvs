package redisWeb

import (
	"net/http"

	"github.com/CunningMatthew/redisutil"
	"github.com/gorilla/mux"
)

func setHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]
	value := vars["value"]
	message, err := redisUtil.SetString(key, value)
	if err != nil {
		http.Error(writer, message, 400)
	} else {
		writer.WriteHeader(200)
		writer.Write([]byte(message))
	}
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["key"]

	message, err := redisUtil.GetString(key)
	if err != nil {
		http.Error(writer, message, 400)
	} else {
		writer.Write([]byte(message))
	}
}

//RegisterRoutes will register the redis web endpoints
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/add/{key}/{value}", setHandler).Methods("GET")
	router.HandleFunc("/add/{key}", getHandler).Methods("GET")
}
