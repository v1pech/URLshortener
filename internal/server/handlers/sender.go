package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func sendResponse(w http.ResponseWriter, status int, result string, err error) {
	response := Response{
		Status: strconv.Itoa(status),
		Result: result,
	}
	if err != nil {
		response.Error = err.Error()
	}
	jsonData, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(jsonData)
}
