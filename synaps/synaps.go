package synaps

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// This function receives a string named alias which is used as a user id
// and it calls synaps API to initiate a new session, it then returns a session id
func GetSessionId(alias string) (string, error) {

	// We usually use a User Id as an alias e.g. email address
	apiPath := fmt.Sprintf(API_SESSION_INIT_WITH_ALIAS, alias)

	headers := map[string]string{
		"Content-Length": "0",
	}

	data, err := callSynapsAPI(apiPath, REQUEST_METHOD_POST, nil, headers)
	if err != nil {
		return "", err
	}

	var sInfo SessionInfo
	if err := json.Unmarshal(data, &sInfo); err != nil {
		return "", err
	}

	return sInfo.SessionId, nil
}

// This function retrieves the finished sessions
// i.e. Information of the users who have their verification process finalized
func GetFinishedSessions() ([]SessionInfo, error) {
	data, err := callSynapsAPI(API_SESSION_FINISHED, REQUEST_METHOD_GET, nil)
	if err != nil {
		return nil, err
	}

	var sInfoList []SessionInfo
	if err := json.Unmarshal(data, &sInfoList); err != nil {
		return nil, err
	}

	return sInfoList, nil
}

// This function retrieves the sessions which have an ongoing verification process
func GetPendingSessions() ([]SessionInfo, error) {
	data, err := callSynapsAPI(API_SESSION_PENDING, REQUEST_METHOD_GET, nil)
	if err != nil {
		return nil, err
	}

	var sInfoList []SessionInfo
	if err := json.Unmarshal(data, &sInfoList); err != nil {
		return nil, err
	}

	return sInfoList, nil
}

func GetSessionDetails(sessionId string) (SessionDetails, error) {

	headers := map[string]string{
		API_HEADER_SESSION_ID: sessionId,
	}

	data, err := callSynapsAPI(API_SESSION_DETAILS, REQUEST_METHOD_GET, nil, headers)
	if err != nil {
		return SessionDetails{}, err
	}

	var sDetails SessionDetails
	if err := json.Unmarshal(data, &sDetails); err != nil {
		return SessionDetails{}, err
	}

	return sDetails, nil
}

// This function calls the given synaps API
// it receives the api name, method, postBody and can receive some additional headers
// it then receives theresponse body in bytes
func callSynapsAPI(api string, reqMethod string, postBody []byte, additionalHeaders ...map[string]string) ([]byte, error) {

	apiKey := os.Getenv("API_KEY")
	clientId := os.Getenv("CLIENT_ID")

	fullApiPath := os.Getenv("API_PATH") + api

	req, err := http.NewRequest(reqMethod, fullApiPath, bytes.NewBuffer(postBody))
	if err != nil {
		log.Printf("[callSynapsAPI] could not make the request: %v", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(API_HEADER_API_KEY, apiKey)
	req.Header.Set(API_HEADER_CLIENT_ID, clientId)

	if len(additionalHeaders) > 0 {
		for headerName := range additionalHeaders[0] {
			req.Header.Set(headerName, additionalHeaders[0][headerName])
		}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[callSynapsAPI] did not receive a response from Synaps Server: %v", err)
		return nil, err
	}

	if resp.StatusCode != 200 {
		err := fmt.Errorf("synaps api  error (%v): %v \n\tapi path: %v", resp.StatusCode, resp.Status, fullApiPath)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
