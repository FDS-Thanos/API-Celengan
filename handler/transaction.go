package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	GetTransaction(*gin.Context)
	Create(*gin.Context)
}
type transactionImplement struct{}

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

type ListBankInterface interface {
	ReceiveBank(*gin.Context)
}

type ListBankImplement struct{}

func (i *ListBankImplement) ReceiveBank(g *gin.Context) {

	// Perform HTTP request to external service
	data, err := fetchDataListBank()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, data)
}

type TransferListBank struct {
	Data []struct {
		Name   string
		BankID string
	}
}

func fetchDataListBank() (*TransferListBank, error) {
	var client = &http.Client{}
	var data TransferListBank
	var err error

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/transfer/list-bank", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type ListAccountInterface interface {
	ReceiveAccount(*gin.Context)
}

type ListAccountImplement struct{}

func (a *ListAccountImplement) ReceiveAccount(g *gin.Context) {

	// Perform HTTP request to external service
	data, err := fetchDataListAccount()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, data)
}

type GetListAccount struct {
	Data struct {
		Name      string
		BankID    string
		AccountID string
	}
}

func fetchDataListAccount() (*GetListAccount, error) {
	var client = &http.Client{}
	var data GetListAccount
	var err error

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/transfer/list-account", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type ValidAccountInterface interface {
	ValidAccount(*gin.Context)
}

type ValidAccountImplement struct {
}

func (v *ValidAccountImplement) ValidAccount(g *gin.Context) {
	bankid := g.Param("bankid")
	accountid := g.Param("accountid")
	// Perform HTTP request to external service
	data, err := fetchDataCheckValidAccount(bankid, accountid)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, data)
}

type CheckValidAccount struct {
	Account struct {
		BankID    string
		AccountID string
		Name      string
	}
}

func fetchDataCheckValidAccount(bankid string, accountid string) (*CheckValidAccount, error) {
	var client = &http.Client{}
	var data CheckValidAccount
	var err error

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/transfer/"+bankid+"/"+accountid, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

type TransferAmountInterface interface {
	SendAmount(*gin.Context)
}

type TransferAmountImplement struct{}

func (i *TransferAmountImplement) SendAmount(g *gin.Context) {

	// Perform HTTP request to external service
	data, err := fetchDataTransferAmount()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, data)
}

type TransferAmount struct {
	Data []struct {
		AccountID string
		BankID    string
		Amount    string
	}
}

func fetchDataTransferAmount() (*TransferAmount, error) {
	var client = &http.Client{}
	var data TransferAmount
	var err error

	request, err := http.NewRequest("POST", "http://localhost:5000/v1/api/transfer/", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

// func (a *transferamountImplement) CreateAccount(g *gin.Context) {
// 	BodyPayLoad := model.Account{}

// 	err := g.BindJSON(&BodyPayLoad)
// 	if err != nil {
// 		g.AbortWithStatusJSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	orm := utils.NewDatabase().Orm
// 	db, _ := orm.DB()

// 	defer db.Close()

// 	result := orm.Create(&BodyPayLoad)
// 	if result.Error != nil {
// 		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
// 			"error": result.Error,
// 		})
// 		return
// 	}

// 	g.JSON(http.StatusOK, gin.H{
// 		"message": "Transfer Successfully",
// 		"data":    BodyPayLoad,
// 	})

// }

type TransactionHistoryInterface interface {
	ReceiveTransaction(*gin.Context)
}

type TransactionHistoryImplement struct{}

func (i *TransactionHistoryImplement) ReceiveTransaction(g *gin.Context) {

	// Perform HTTP request to external service
	data, err := fetchTransactionHistory()
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	g.JSON(http.StatusOK, data)
}

type TransactionHistory struct {
	Data []struct {
		Name   string
		BankID string
	}
}

func fetchTransactionHistory() (*TransactionHistory, error) {
	var client = &http.Client{}
	var data TransactionHistory
	var err error

	request, err := http.NewRequest("GET", "http://localhost:5000/v1/api/transfer/", nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
