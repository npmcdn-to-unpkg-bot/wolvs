package main

import (
	"flag"
	"net/http"

	"github.com/CunningMatthew/wolvs/redis/redisUtil"
	"github.com/CunningMatthew/wolvs/redisweb"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Gorilla!\n"))
}

func main() {
	flag.Parse()
	redisUtil.Initialize()
	defer redisUtil.Close()

	router := mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods("GET")
	redisWeb.RegisterRoutes(router)

	n := negroni.Classic()

	n.UseHandler(router)
	n.Run(":8080")
}
