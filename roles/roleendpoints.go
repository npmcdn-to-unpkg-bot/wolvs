package roles

import (
	"net/http"

	"github.com/CunningMatthew/wolvs/common/web"
	"github.com/CunningMatthew/wolvs/redis/redisUtil"
	"github.com/gorilla/mux"
)

func getHandler(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	message, err := redisUtil.GetString(vars["key"])
	commonWeb.WriteMessageOrError(writer, message, err)
}

//RegisterRoutes will register the redis web endpoints
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/roles", getHandler).Methods("GET")
}
