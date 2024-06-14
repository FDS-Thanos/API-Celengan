package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AccountInterface interface {
	GetAccount(*gin.Context)
	CreateAccount(*gin.Context)
	UpdateAccount(*gin.Context)
	DeleteAccount(*gin.Context)
	BalanceAccount(*gin.Context)
}

type acccountImplement struct{}

func NewAccount() AccountInterface {
	return &acccountImplement{}
}

func (a *acccountImplement) GetAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	name := queryParam.Get("name")

	accounts := []model.Account{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	q := orm
	if name != "" {
		q = q.Where("name = ?", name)
	}

	result := q.Find(&accounts)

	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    accounts,
	})
}

func (a *acccountImplement) CreateAccount(g *gin.Context) {
	BodyPayLoad := model.Account{}

	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

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
		"message": "Create account successfully :)",
		"data":    BodyPayLoad,
	})

}

func (a *acccountImplement) UpdateAccount(g *gin.Context) {
	BodyPayLoad := model.Account{}
	err := g.BindJSON(&BodyPayLoad)
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	user := model.Account{}
	orm.First(&user, "account_id = ?", id)
	user.Name = BodyPayLoad.Name
	user.Username = BodyPayLoad.Username
	orm.Save(&user)

	if user.AccountID == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "Data not found :(",
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Update account successfully :)",
		"data":    user,
	})
}

func (a *acccountImplement) DeleteAccount(g *gin.Context) {
	id := g.Param("id")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Where("account_id = ?", id).Delete(&model.Account{})
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	g.JSON(http.StatusOK, gin.H{
		"message": "Delete Account successfully",
		"data":    id,
	})

}

func (a *acccountImplement) BalanceAccount(g *gin.Context) {
	queryParam := g.Request.URL.Query()

	balance := queryParam.Get("balance")

	g.JSON(http.StatusOK, gin.H{
		"message": "Get account successfully",
		"data":    map[string]string{"name": balance},
	})

}
