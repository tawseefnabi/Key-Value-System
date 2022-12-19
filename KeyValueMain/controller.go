package main

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	service *Service
}

func NewController(service *Service) *Controller {
	return &Controller{
		service: service,
	}
}
func (c *Controller) put(w http.ResponseWriter, r *http.Request) {
	var keyValuePair KeyValuePair
	err := json.NewDecoder(r.Body).Decode(&keyValuePair)
	if err != nil {
		CommonHttpRequest(w, http.StatusBadRequest, err.Error(), "err")
		return
	}
	if keyValuePair.Ttl < 0 {
		CommonHttpRequest(w, http.StatusBadRequest, "ttl cannot be negative", "err")
		return
	}
	err = c.service.put(keyValuePair)
	if err != nil {
		CommonHttpRequest(w, http.StatusBadRequest, err.Error(), "err")
		return
	}
	CommonHttpRequest(w, http.StatusOK, "created", "normal")

}

func (c *Controller) get(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("key")
	if key == "" {
		CommonHttpRequest(w, http.StatusBadRequest, "plz provide key", "err")
		return
	}
	value, err := c.service.get(key)
	if err != nil {
		CommonHttpRequest(w, http.StatusInternalServerError, err.Error(), "err")
		return
	}
	CommonHttpRequest(w, http.StatusOK, "value:="+value, "normal")

}
