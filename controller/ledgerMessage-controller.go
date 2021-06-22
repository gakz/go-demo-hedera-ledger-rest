package controller

import (
	"github.com/droatl2000/demo-ledger-rest/model"
	"github.com/droatl2000/demo-ledger-rest/service"
	"github.com/gin-gonic/gin"
)

type LedgerMessageController interface {
	FindByVin(vin string) []model.LedgerMessage
	FindAll() []model.LedgerMessage
	Save(ctx *gin.Context) model.LedgerMessage
}

type ledgerMessageController struct {
	service service.LedgerMessageService
}

func New(service service.LedgerMessageService) LedgerMessageController {
	return &ledgerMessageController{
		service: service,
	}
}

func (c *ledgerMessageController) FindAll() []model.LedgerMessage {
	return c.service.FindAll()
}

func (c *ledgerMessageController) Save(ctx *gin.Context) model.LedgerMessage {
	var ledgerMessage model.LedgerMessage

	ctx.BindJSON(&ledgerMessage)

	c.service.Save(ledgerMessage)
	return ledgerMessage
}

func (c *ledgerMessageController) FindByVin(vin string) []model.LedgerMessage {
	return c.service.FindByVin(vin)
}
