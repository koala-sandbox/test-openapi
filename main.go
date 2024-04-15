package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	api "github.com/koala-sandbox/test-openapi/api"
)

func configureHTTPAPI() *mux.Router {
	ItemApiService := api.NewItemApiService()
	ItemApiController := api.NewItemApiController(ItemApiService)
	KoalaApiService := api.NewKoalaApiService()
	KoalaApiController := api.NewKoalaApiController(KoalaApiService)
	router := api.NewRouter(ItemApiController, KoalaApiController)
	return router
}

// defineFlags contains all the flag definitions, which are called before flag.Parse().
func parseFlags() {
	flag.Parse()
}

func main() {
	log.Printf("Server started.")
	parseFlags()

	router := configureHTTPAPI()
	router.HandleFunc("/koala", HelloKoala)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Fatal(srv.ListenAndServe())
}

func HelloKoala(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	VERSION, err := ioutil.ReadFile("VERSION")
	if err != nil {
		fmt.Fprintln(w, "Error reading VERSION file (test-openapi): "+err.Error())
		return
	}
	fmt.Fprintln(w, "Hello world! service:test-openapi version: "+string(VERSION))
}
