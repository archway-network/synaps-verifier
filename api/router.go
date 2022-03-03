package api

import (
	"log"
	"net/http"

	routing "github.com/julienschmidt/httprouter"
)

/*-------------------------*/

func setupRouter() *routing.Router {

	var router = routing.New()

	router.GET("/", IndexPage)
	router.GET("/ui/*file_path", UI)

	router.GET("/sessionid/:email", GetSynapsSessionId)

	return router
}

/*-------------------------*/

// ListenAndServeHTTP serves the APIs and the ui
func ListenAndServeHTTP(addr string) {

	router := setupRouter()
	if addr == "" {
		addr = ":80"
	}

	log.Printf("[INFO ] Serving on %s", addr)

	log.Fatal(http.ListenAndServe(addr, router))
}

/*-------------------------*/
