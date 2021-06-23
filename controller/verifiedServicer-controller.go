package controller

import (
	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-gonic/gin"
)

type VerifiedServicerController interface {
	FindAll() []model.VerifiedServicer
	Save(ctx *gin.Context) model.VerifiedServicer
}

type verifiedServicerController struct {
	service service.VerifiedServicerService
}

func NewServicer(service service.VerifiedServicerService) VerifiedServicerController {
	return &verifiedServicerController{
		service: service,
	}
}

func (c *verifiedServicerController) FindAll() []model.VerifiedServicer {
	return c.service.FindAll()
}

func (c *verifiedServicerController) Save(ctx *gin.Context) model.VerifiedServicer {
	var verifiedServicer model.VerifiedServicer

	ctx.BindJSON(&verifiedServicer)

	c.service.Save(verifiedServicer)
	return verifiedServicer
}
