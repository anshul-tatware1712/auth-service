package utils

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validator *validator.Validate

func init(){
	Validator = validator.New(validator.WithRequiredStructEnabled())
}

func WriteJsonResponse(w http.ResponseWriter,status int,data any) error{
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func WriteJsonToSuccessResponse(w http.ResponseWriter,data any) error{
	response:=map[string]any{
		"error":"nil",
		"message":"ok",
		"status":http.StatusOK,
		"data":data,
	}
	return WriteJsonResponse(w,http.StatusOK,response)
}

func WriteJsonToErrorResponse(w http.ResponseWriter,status int,err string) error{
	response:=map[string]any{
		"error":err,
		"status":status,
	}
	return WriteJsonResponse(w,status,response)
}

func ReadJsonRequest(r *http.Request, v any) error {
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	return decoder.Decode(v)
}