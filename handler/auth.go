package handler

import (
	"api_gateway/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	AuthLogin(*gin.Context)
}

type AuthImplement struct{}

func Login() AuthInterface {
	return &AuthImplement{}
}

type BodyPayLoadAuth struct {
	Username string
	Password string
}

func (a *AuthImplement) AuthLogin(c *gin.Context) {
	bodyPayloadAuth := BodyPayLoadAuth{}

	err := c.BindJSON(&bodyPayloadAuth)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	authenticated := usecase.NewLogin().Autentikasi(bodyPayloadAuth.Username, bodyPayloadAuth.Password)

	if authenticated {
		c.JSON(http.StatusOK, gin.H{
			"message": "Anda berhasil login",
			"data":    bodyPayloadAuth,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Anda gagal login",
			"data":    err,
		})
	}
}
