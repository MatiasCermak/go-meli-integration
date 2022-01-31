package controller

import (
	"net/http"
	"strconv"
	"sync"

	"github.com/MatiasCermak/go-meli-integration/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ItemsAll(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With, Access-Control-Allow-Origin")

	if c.Request.Method == "OPTIONS" {
		c.Writer.WriteHeader(http.StatusOK)
		return
	}

	token := c.Query("token")
	userid := c.Query("userid")
	var url string = "https://api.mercadolibre.com/users/" + userid + "/items/search?access_token=" + token
	var res ItemIds
	getAndMarshall(url, &res, c)
	if len(res.Results) == 0 {
		c.JSON(400, struct {
			Error string
		}{
			Error: "No items being sold by this user",
		})
		return
	}
	var itmCar FinalType
	DB := db.Initialize()
	itmCar.m = make(map[string]ItemCarrierDefinite)
	ch3 := make(chan []SalesCarrier)
	go salesCollector(token, userid, ch3, c)
	var wg sync.WaitGroup
	for i := 0; i < len(res.Results); i++ {
		wg.Add(2)
		go itemCollector(token, res.Results[i], &itmCar, c, &wg, DB)
		go questionCollector(token, res.Results[i], &itmCar, c, &wg)
	}
	var finalResp ResponseCarrier
	wg.Wait()
	values := []ItemCarrierDefinite{}
	for _, value := range itmCar.m {
		values = append(values, value)
	}
	finalResp.Items = values
	finalResp.Sales = <-ch3
	if err := DB.Save(&finalResp.Sales).Error; err != nil {
		// always handle error like this, cause errors maybe happened when connection failed or something.
		// record not found...
		if gorm.ErrRecordNotFound == err {
			DB.Create(&finalResp.Sales) // create new record from newUser
		}
	}
	c.JSON(200, finalResp)
}

func itemCollector(token, itemid string, itmcar *FinalType, c *gin.Context, wg *sync.WaitGroup, DB *gorm.DB) {
	defer wg.Done()
	var url string = "https://api.mercadolibre.com/items?ids=" + itemid + "&attributes=id,price,available_quantity,title,pictures,sold_quantity&access_token=" + token
	var res []Items
	getAndMarshall(url, &res, c)
	var resp ItemCarrier
	resp.Quantity = res[0].Body.Available_quantity
	resp.SoldQuantity = res[0].Body.Sold_quantity
	resp.Id, _ = strconv.ParseInt(res[0].Body.Id[3:], 10, 64)
	resp.Picture = res[0].Body.Pictures[0]["url"]
	resp.Price = res[0].Body.Price
	resp.Title = res[0].Body.Title
	itmcar.RWMutex.Lock()
	if entry, ok := itmcar.m[itemid]; ok {
		entry.Item = resp
		itmcar.m[itemid] = entry
	} else {
		itmcar.m[itemid] = ItemCarrierDefinite{Item: resp, Questn: Questions{}}
	}
	if err := DB.Save(itmcar.m[itemid].Item).Error; err != nil {
		// always handle error like this, cause errors maybe happened when connection failed or something.
		// record not found...
		if gorm.ErrRecordNotFound == err {
			DB.Create(itmcar.m[itemid].Item) // create new record from newUser
		}
	}
	itmcar.RWMutex.Unlock()
}

func questionCollector(token, itemid string, itmcar *FinalType, c *gin.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	var url string = "https://api.mercadolibre.com/questions/search?item=" + itemid + "&access_token=" + token
	var res Question
	getAndMarshall(url, &res, c)
	var unansweredQ Question
	for i := len(res.Questn) - 1; i >= 0; i-- {
		if res.Questn[i].Status == "UNANSWERED" {
			itmcar.Lock()
			if entry, ok := itmcar.m[itemid]; ok {
				entry.Questn = append(entry.Questn, res.Questn[i])
				itmcar.m[itemid] = entry
			} else {
				itmcar.m[itemid] = ItemCarrierDefinite{Item: ItemCarrier{}, Questn: append(unansweredQ.Questn, res.Questn[i])}
			}
			itmcar.Unlock()
		}
	}
}

func salesCollector(token, userid string, ch3 chan []SalesCarrier, c *gin.Context) {
	var url string = "https://api.mercadolibre.com/orders/search?seller=" + userid + "&order.status=paid&access_token=" + token
	var res Sales
	var resp []SalesCarrier
	getAndMarshall(url, &res, c)
	if len(res.Results) == 0 {
		resp = []SalesCarrier{}
	}
	for i := 0; i < len(res.Results); i++ {
		for j := 0; j < len(res.Results[i].Payments); j++ {
			var sales SalesCarrier
			sales.Id = res.Results[i].Payments[j].ID
			sales.Title = res.Results[i].Payments[j].Reason
			sales.Date = res.Results[i].Payments[j].DateApproved
			sales.PriceTotal = res.Results[i].Payments[j].TotalPaidAmount
			sales.Price = res.Results[i].Payments[j].TransactionAmount
			resp = append(resp, sales)
		}

	}

	ch3 <- resp
}
