package myflyingbox

import "errors"

// Response is the base response type for all API responses.
type Response struct {
	Status    string      `json:"status"`
	Self      string      `json:"self"`
	Count     int         `json:"count"`
	Data      interface{} `json:"data"`
	ErrorInfo *ErrorInfo  `json:"error"`
}

func (r Response) Error() error {
	if r.Status == "failure" {
		if r.ErrorInfo != nil {
			return r.ErrorInfo.Error()
		}
		return errors.New(r.Status)
	}
	return nil
}

// ErrorInfo is an upstream error
type ErrorInfo struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

func (e ErrorInfo) Error() error {
	return errors.New(e.Message)
}
