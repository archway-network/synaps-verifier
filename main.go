package main

import (
	"os"

	"github.com/archway-network/synaps-verifier/api"
)

/*--------------*/

func main() {

	api.ListenAndServeHTTP(os.Getenv("SERVING_ADDR"))
}
