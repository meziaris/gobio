package handler

import (
	"encoding/json"
	"gobio/internal/pkg/validator"
	"net/http"
)

// Handle response error
func ResponseError(w http.ResponseWriter, statusCode int, message string) {
	resp := ResponseBody{
		Status:  "error",
		Message: message,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
}

// Handle response success
func ResponseSuccess(w http.ResponseWriter, statusCode int, message string, data interface{}) {
	resp := ResponseBody{
		Status:  "success",
		Message: message,
		Data:    data,
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(resp)
}

// Parse request data & validate struct
func BindAndCheck(w http.ResponseWriter, r *http.Request, data interface{}) bool {
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(data)
	if err != nil {
		ResponseError(w, http.StatusUnprocessableEntity, err.Error())
		return true
	}

	isError := validator.Check(data)
	if isError {
		ResponseError(w, http.StatusUnprocessableEntity, "request format is not valid")
		return true
	}

	return false
}
