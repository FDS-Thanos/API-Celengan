package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	GetTransaction(*gin.Context)
	Create(*gin.Context)
}
type transactionImplement struct{}

// GetTransaction implements TransactionInterface.
func (pi *transactionImplement) GetTransaction(*gin.Context) {

}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

func (pi *transactionImplement) Create(g *gin.Context) {
	BodyPayLoad := model.Transaction{}

	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	timeNow := time.Now()
	BodyPayLoad.Transaction_date = &timeNow

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Create(&BodyPayLoad)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Create transaction successfully :D",
		"data":    BodyPayLoad,
	})
}
