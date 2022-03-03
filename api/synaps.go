package api

import (
	"log"
	"net/http"

	"github.com/archway-network/synaps-verifier/synaps"
	"github.com/archway-network/synaps-verifier/tools"
	routing "github.com/julienschmidt/httprouter"
)

/*----------------------*/

/*
* This function implements GET /sessionid/:email
 */
func GetSynapsSessionId(resp http.ResponseWriter, req *http.Request, params routing.Params) {

	email := params.ByName("email")
	sessionId, err := synaps.GetSessionId(email)

	if err != nil {
		log.Printf("Error in synaps.GetSessionId : %v", err)
		http.Error(resp, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tools.SendJSON(resp, sessionId)
}

/*-------------*/
