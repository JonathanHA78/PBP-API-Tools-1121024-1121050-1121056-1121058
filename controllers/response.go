package controllers

import (
	"encoding/json"
	_ "fmt"
	"net/http"
)

func Response(w http.ResponseWriter, req interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(req)
}

func SendSuccessResponse(w http.ResponseWriter, message string) {
	var response SuccessResponse
	response.Status = 200
	response.Message = message
	Response(w, response)
}

func SendErrorResponse(w http.ResponseWriter, message string) {
	var response ErrorResponse
	response.Status = 400
	response.Message = message
	Response(w, response)
}

func sendUnAuthorizedResponse(w http.ResponseWriter) {
	var response ErrorResponse
	response.Status = 401
	response.Message = "Unauthorized Access"
	Response(w, response)
}
