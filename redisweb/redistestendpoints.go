package redisWeb

import (
	"net/http"

	"github.com/CunningMatthew/wolvs/common/web"
	"github.com/CunningMatthew/wolvs/redis/redisUtil"
	"github.com/gorilla/mux"
)

func setHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	message, err := redisUtil.SetString(vars["key"], vars["value"])
	commonWeb.WriteAcceptedMessageOrError(writer, message, err)
}

func getHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	message, err := redisUtil.GetString(vars["key"])
	commonWeb.WriteMessageOrError(writer, message, err)
}

//RegisterRoutes will register the redis web endpoints
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/add/{key}/{value}", setHandler).Methods("GET")
	router.HandleFunc("/add/{key}", getHandler).Methods("GET")
}
