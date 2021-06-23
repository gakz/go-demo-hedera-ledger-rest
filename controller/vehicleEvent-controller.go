package controller

import (
	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-gonic/gin"
)

type VehicleEventController interface {
	FindByVin(vin string) []model.VehicleEvent
	FindAll() []model.VehicleEvent
	Save(ctx *gin.Context) model.VehicleEvent
}

type vehicleEventController struct {
	service service.VehicleEventService
}

func NewEvent(service service.VehicleEventService) VehicleEventController {
	return &vehicleEventController{
		service: service,
	}
}

func (c *vehicleEventController) FindAll() []model.VehicleEvent {
	return c.service.FindAll()
}

func (c *vehicleEventController) Save(ctx *gin.Context) model.VehicleEvent {
	var vehicleEvent model.VehicleEvent

	ctx.BindJSON(&vehicleEvent)

	c.service.Save(vehicleEvent)
	return vehicleEvent
}

func (c *vehicleEventController) FindByVin(vin string) []model.VehicleEvent {
	return c.service.FindByVin(vin)
}
