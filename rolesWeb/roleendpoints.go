package rolesWeb

import (
	"encoding/json"
	"net/http"

	"github.com/CunningMatthew/wolvs/common"
	"github.com/CunningMatthew/wolvs/common/web"
	"github.com/gorilla/mux"
)

func getHandler(writer http.ResponseWriter, request *http.Request) {
	roles := &[]common.Role{
		common.Role{Name: "Villager", Description: "Lowly Villager", MaximumRatio: 8},
		common.Role{Name: "Werewolv", Description: "Silly Super Innocent Puppy", MaximumRatio: 3},
	}
	message, err := json.Marshal(roles)
	commonWeb.WriteMessageOrError(writer, string(message), err)
}

//RegisterRoutes will register the redis web endpoints
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/roles", getHandler).Methods("GET")
}
