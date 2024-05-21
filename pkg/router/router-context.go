package router

import (
	"encoding/json"
	"errors"
	"net/http"

	"xcylla.io/config/pkg/utils"
)

type RouterContext struct {
	http.Handler
	Request *http.Request
	Writer  http.ResponseWriter
}

type finalResponse struct {
	Error *string `json:"error"`
	Data  any     `json:"data"`
}

// BindBody decodes the JSON body into the provided struct
func (rc *RouterContext) BindBody(reqData any) error {
	if rc.Request.Body == nil {
		return errors.New("request body is empty")
	}
	err := utils.Decode(rc.Request.Body, &reqData)
	if err != nil {
		logging.Error("Error decoding request body: %v", err)
		return err
	}
	return nil
}

// returnWithCode constructs a response with a given HTTP status code
func (rc *RouterContext) returnWithCode(code int, data any, err error) {
	finalResponse := &finalResponse{
		Error: getErrorString(err),
		Data:  data,
	}

	respData, pErr := json.Marshal(finalResponse)
	if pErr != nil {
		logging.Error("Error parsing response: %v", pErr)
		rc.Writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	rc.Writer.Header().Set("Content-Type", "application/json")
	rc.Writer.WriteHeader(code)
	_, writeErr := rc.Writer.Write(respData)
	if writeErr != nil {
		logging.Error("Error writing response: %v", writeErr)
	}
}

func (rc *RouterContext) ReturnBadRequest(data any, err error) {
	rc.returnWithCode(http.StatusBadRequest, data, err)
}

func (rc *RouterContext) ReturnInternalError(data any, err error) {
	rc.returnWithCode(http.StatusInternalServerError, data, err)
}

func (rc *RouterContext) ReturnOK(data any) {
	rc.returnWithCode(http.StatusOK, data, nil)
}

func (rc *RouterContext) AuthenticateToken() bool {
	token := rc.Request.Header.Get("Authorization")
	if token == "" {
		rc.returnWithCode(http.StatusForbidden, nil, errors.New("token not found"))
		return false
	}

	// isValid := actionsAuth.IsTokenValid(token)
	// if !isValid {
	// 	rc.returnWithCode(http.StatusForbidden, nil, errors.New("invalid token"))
	// 	return false
	// }

	return true
}

// getErrorString returns the error message as a string pointer, or nil if no error
func getErrorString(err error) *string {
	if err != nil {
		str := err.Error()
		return &str
	}
	return nil
}
