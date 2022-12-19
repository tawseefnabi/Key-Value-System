package controller

import (
	service "KeyValueMain/Service"
)

type Controller struct {
	service *service.Service
}

func NewController(ser *service.Service) *Controller {
	return &Controller{
		service: ser,
	}
}
