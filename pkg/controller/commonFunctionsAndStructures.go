package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

//-Common Structures------------------------------------

type ItemIds struct {
	Results []string `json:"results"`
}

type Items struct {
	Body struct {
		Id                 string
		Title              string
		Price              float32
		Pictures           []map[string]string
		Available_quantity int
		Sold_quantity      int
	}
}

type Questions []struct {
	DateCreated        string `json:"date_created"`
	ItemID             string `json:"item_id"`
	SellerID           int    `json:"seller_id"`
	Status             string `json:"status"`
	Text               string `json:"text"`
	ID                 int64  `json:"id" gorm: primaryKey`
	DeletedFromListing bool   `json:"deleted_from_listing"`
	//Answer                string `json:"answer"`
	ItemCarrierDefiniteID int64
}

type Question struct {
	Questn Questions `json:"questions"`
}

type Sales struct {
	Results []struct {
		Payments []struct {
			Reason            string      `json:"reason"`
			StatusCode        interface{} `json:"status_code"`
			TotalPaidAmount   float64     `json:"total_paid_amount"`
			OperationType     string      `json:"operation_type"`
			TransactionAmount float64     `json:"transaction_amount"`
			DateApproved      string      `json:"date_approved"`
			Collector         struct {
				ID int `json:"id"`
			} `json:"collector"`
			CouponID             interface{} `json:"coupon_id"`
			Installments         int         `json:"installments"`
			AuthorizationCode    string      `json:"authorization_code"`
			TaxesAmount          int         `json:"taxes_amount"`
			ID                   int64       `json:"id"`
			DateLastModified     string      `json:"date_last_modified"`
			CouponAmount         int         `json:"coupon_amount"`
			AvailableActions     []string    `json:"available_actions"`
			ShippingCost         float64     `json:"shipping_cost"`
			InstallmentAmount    float64     `json:"installment_amount"`
			DateCreated          string      `json:"date_created"`
			ActivationURI        interface{} `json:"activation_uri"`
			OverpaidAmount       int         `json:"overpaid_amount"`
			CardID               int         `json:"card_id"`
			StatusDetail         string      `json:"status_detail"`
			IssuerID             string      `json:"issuer_id"`
			PaymentMethodID      string      `json:"payment_method_id"`
			PaymentType          string      `json:"payment_type"`
			DeferredPeriod       interface{} `json:"deferred_period"`
			AtmTransferReference struct {
				TransactionID string      `json:"transaction_id"`
				CompanyID     interface{} `json:"company_id"`
			} `json:"atm_transfer_reference"`
			SiteID             string      `json:"site_id"`
			PayerID            int         `json:"payer_id"`
			MarketplaceFee     float64     `json:"marketplace_fee"`
			OrderID            int         `json:"order_id"`
			CurrencyID         string      `json:"currency_id"`
			Status             string      `json:"status"`
			TransactionOrderID interface{} `json:"transaction_order_id"`
		} `json:"payments"`
	}
}

type ItemCarrier struct {
	Id           int64
	Title        string
	Price        float32
	Quantity     int
	SoldQuantity int
	Picture      string
}

type ItemCarrierDefinite struct {
	Item ItemCarrier

	Questn Questions
}

type SalesCarrier struct {
	Id         int64
	Title      string
	Date       string
	Price      float64
	PriceTotal float64
}

type ResponseCarrier struct {
	Items []ItemCarrierDefinite
	Sales []SalesCarrier
}

type FinalType struct {
	sync.RWMutex
	m map[string]ItemCarrierDefinite
}

type AnswerOut map[string]interface{}

//-Common Functions----------------------------------------

func getAndMarshall(url string, res interface{}, c *gin.Context) {
	req, erro := http.Get(url)
	if req.StatusCode != 200 || erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}
	defer req.Body.Close()
	data, erro := ioutil.ReadAll(req.Body)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error())
		return
	}

	erro = json.Unmarshal(data, &res)
	if erro != nil {
		c.JSON(req.StatusCode, erro.Error()+url)
		return
	}
}
