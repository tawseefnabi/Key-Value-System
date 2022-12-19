package main

import (
	"encoding/json"
	"net/http"
)

func CommonHttpRequest(w http.ResponseWriter, statusCode int, message string, reqType string) {
	w.WriteHeader(statusCode)
	w.Header().Add("Content-Type", "application/json")
	if reqType == "err" {
		msg := map[string]interface{}{
			"status":  "error",
			"message": message,
		}
		data, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Write(data)
	} else {
		msg := map[string]interface{}{
			"status":  "ok",
			"message": message,
		}
		data, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		w.Write(data)
	}
}
