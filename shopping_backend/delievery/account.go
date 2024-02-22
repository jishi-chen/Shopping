package delievery

import (
	"shopping_backend/domain"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// AccountHandler ...
type AccountHandler struct {
	AccountUsecase   domain.AccountUsecase
	CommodityUsecase domain.CommodityUsecase
}

// NewAccountHandler ...
func NewAccountHandler(e *gin.Engine, accountUsecase domain.AccountUsecase, commodityUsecase domain.CommodityUsecase) {
	handler := &AccountHandler{
		AccountUsecase:   accountUsecase,
		CommodityUsecase: commodityUsecase,
	}

	e.GET("/api/v1/accounts/:accountID", handler.GetAccountByAccountID)
	e.POST("/api/v1/accounts", handler.PostToCreateAccount)
	e.POST("/api/v1/accounts/:accountID/foster", handler.PostToFosterAccount)
}

// GetAccountByAccountID ...
func (d *AccountHandler) GetAccountByAccountID(c *gin.Context) {
	accountID := c.Param("accountID")

	anAccount, err := d.AccountUsecase.GetByID(c, accountID)
	if err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Query account error",
		})
		return
	}

	c.JSON(200, &swagger.AccountInfo{
		Id:     anAccount.ID,
		Name:   anAccount.Name,
		Status: anAccount.Status,
	})
}

// PostToCreateAccount ...
func (d *AccountHandler) PostToCreateAccount(c *gin.Context) {
	var body swagger.AccountInfoRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}
	aAccount := domain.Account{
		Name: body.Name,
	}
	if err := d.AccountUsecase.Store(c, &aAccount); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	c.JSON(200, swagger.AccountInfo{
		Id:     aAccount.ID,
		Name:   aAccount.Name,
		Status: aAccount.Status,
	})
}

// PostToFosterAccount ...
func (d *AccountHandler) PostToFosterAccount(c *gin.Context) {
	accountID := c.Param("accountID")

	var body swagger.FosterRequest
	if err := c.BindJSON(&body); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Parsing failed",
		})
		return
	}

	if err := d.CommodityUsecase.Store(c, &domain.Commodity{
		UserID: accountID,
		Name:   body.Food.Name,
	}); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}

	if err := d.AccountUsecase.UpdateStatus(c, &domain.Account{
		ID:     accountID,
		Status: "good",
	}); err != nil {
		logrus.Error(err)
		c.JSON(500, &swagger.ModelError{
			Code:    3000,
			Message: "Internal error. Store failed",
		})
		return
	}
	c.JSON(204, nil)
}
