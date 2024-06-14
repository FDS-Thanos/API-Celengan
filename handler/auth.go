package handler

import (
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
	Username string `json:"username"`
	Password string `json:"password"`
}

func (a *AuthImplement) AuthLogin(c *gin.Context) {
	var bodyPayload BodyPayLoadAuth

	err := c.BindJSON(&bodyPayload)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	if bodyPayload.Username == "admin" && bodyPayload.Password == "admin123" {
		c.JSON(http.StatusOK, gin.H{
			"message": "Account retrieved successfully",
			"data":    bodyPayload,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized: Invalid username or password",
		})
	}
}

// type AuthdbInterface interface {
// 	LoginRequest(*gin.Context)
// }

// type AuthdbImplement struct{}

// func Logindb() AuthdbInterface {
// 	return &AuthdbImplement{}
// }

// type AuthdbImplement struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }

// func (l *AuthdbImplement) Logindb(g *gin.Context) {

// 	// Perform HTTP request to external service
// 	data, err := fetchDataLogin()
// 	if err != nil {
// 		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}
// 	g.JSON(http.StatusOK, data)

// 	if data.Username == "" && data.Password == "" {
// 		c.JSON(http.StatusOK, gin.H{
// 			"message": "Login Account successfully",
// 			"data":    data,
// 		})
// 	} else {
// 		c.JSON(http.StatusUnauthorized, gin.H{
// 			"message": "Invalid username or password",
// 		})
// 	}
// }

// func fetchDataLogin() (*AuthdbImplement, error) {
// 	var client = &http.Client{}
// 	var data AuthdbImplement
// 	var err error

// 	request, err := http.NewRequest("POST", "http://localhost:5000/v1/api/login/", nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response, err := client.Do(request)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer response.Body.Close()

// 	err = json.NewDecoder(response.Body).Decode(&data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &data, nil
// }
