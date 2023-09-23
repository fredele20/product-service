package controllers

import "product-service/services"


type Controllers struct {
	services *services.Service
}

func NewController(s *services.Service) *Controllers {
	return &Controllers{
		services: s,
	}
}