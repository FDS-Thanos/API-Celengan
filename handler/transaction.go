package handler

import (
	"api_gateway/model"
	"api_gateway/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type TransactionInterface interface {
	GetTransaction(*gin.Context)
	Create(*gin.Context)
}
type transactionImplement struct{}

func (t *transactionImplement) GetTransaction(*gin.Context) {

}

func NewTransaction() TransactionInterface {
	return &transactionImplement{}
}

func (t *transactionImplement) Create(g *gin.Context) {
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

func (b *ListBankImplement) ReceiveBank(g *gin.Context) {

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
	CreateTransaction(*gin.Context)
}

type TransferAmountImplement struct{}

type BodyPayloadCreateTrx struct {
	AccountID string
	Amount    float64
	BankID    string
}

func (tf *TransferAmountImplement) CreateTransaction(g *gin.Context) {

	url := "http://localhost:5000/v1/api/transfer"

	dataPayload := BodyPayloadCreateTrx{}

	err := g.BindJSON(&dataPayload)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	byDataPayload, err := json.Marshal(dataPayload)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(byDataPayload))
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	trx := model.Transaction{
		Account_id: dataPayload.AccountID,
		Amount:     int(dataPayload.Amount),
	}

	timeNow := time.Now()
	trx.Transaction_date = &timeNow

	defer db.Close()

	result := orm.Create(&dataPayload)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := io.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	g.JSON(http.StatusOK, gin.H{
		"message": "transfer success",
	})
}

type TransferAmount struct {
	Data []struct {
		Name      string
		AccountID string
		BankID    string
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

type TransactionHistoryInterface interface {
	ReceiveTransaction(*gin.Context)
}

type TransactionHistoryImplement struct {
	ID                    int
	AccountID             string
	Transaction_amount    string
	Transaction_date      string
	Transaction_reference string
}

func (h *TransactionHistoryImplement) ReceiveTransaction(g *gin.Context) {
	transaction := []model.Transaction{}

	orm := utils.NewDatabase().Orm
	db, _ := orm.DB()

	defer db.Close()

	result := orm.Find(&transaction)
	if result.Error != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": result.Error,
		})
		return
	}
	g.JSON(200, transaction)
}
