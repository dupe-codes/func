// Defines a set of helpful functions/types for generating HTTP responses

package web

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Status int
	Error  error
	Data   interface{}
}

// Sends an error response with the given response writer.
// The response status is set to the given status, and the given message
// is sent as the payload of the http response
func SendErrorResponse(resp http.ResponseWriter, errorMsg error, status int) {
	resContent := Response{Status: status, Error: errorMsg}
	response, err := json.MarshalIndent(resContent, "", " ")
	if err != nil {
		http.Error(resp, "Error preparing response", http.StatusInternalServerError)
		return
	}
	http.Error(resp, string(response), status)
	return
}

// Takes the given data and sends it as the payload of an HTTP 200 response
// TODO: Refactor shared code between this and SendErrorResponse
func SendSuccessResponse(resp http.ResponseWriter, data interface{}) {
	resContent := &Response{Status: http.StatusOK, Data: data}
	response, err := json.MarshalIndent(resContent, "", " ")
	if err != nil {
		http.Error(resp, "Error preparing response", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(resp, string(response))
	return
}
