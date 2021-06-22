package controller

import (
	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-gonic/gin"
)

type VehicleActionController interface {
	FindByVin(vin string) []model.VehicleAction
	FindAll() []model.VehicleAction
	Save(ctx *gin.Context) model.VehicleAction
}

type vehicleActionController struct {
	service service.VehicleActionService
}

func New(service service.VehicleActionService) VehicleActionController {
	return &vehicleActionController{
		service: service,
	}
}

func (c *vehicleActionController) FindAll() []model.VehicleAction {
	return c.service.FindAll()
}

func (c *vehicleActionController) Save(ctx *gin.Context) model.VehicleAction {
	var vehicleAction model.VehicleAction

	ctx.BindJSON(&vehicleAction)

	c.service.Save(vehicleAction)
	return vehicleAction
}

func (c *vehicleActionController) FindByVin(vin string) []model.VehicleAction {
	return c.service.FindByVin(vin)
}
