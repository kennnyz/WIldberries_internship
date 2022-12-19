package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/nats-io/nats.go"
	stan "github.com/nats-io/stan.go"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "qwerty2004"
	dbname   = "wbl0"
)

type JsStruct struct {
	OrderUid    string `json:"order_uid"`
	TrackNumber string `json:"track_number"`
	Entry       string `json:"entry"`
	Delivery    struct {
		Name    string `json:"name"`
		Phone   string `json:"phone"`
		Zip     string `json:"zip"`
		City    string `json:"city"`
		Address string `json:"address"`
		Region  string `json:"region"`
		Email   string `json:"email"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction"`
		RequestId    string `json:"request_id"`
		Currency     string `json:"currency"`
		Provider     string `json:"provider"`
		Amount       int    `json:"amount"`
		PaymentDt    int    `json:"payment_dt"`
		Bank         string `json:"bank"`
		DeliveryCost int    `json:"delivery_cost"`
		GoodsTotal   int    `json:"goods_total"`
		CustomFee    int    `json:"custom_fee"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id"`
		TrackNumber string `json:"track_number"`
		Price       int    `json:"price"`
		Rid         string `json:"rid"`
		Name        string `json:"name"`
		Sale        int    `json:"sale"`
		Size        string `json:"size"`
		TotalPrice  int    `json:"total_price"`
		NmId        int    `json:"nm_id"`
		Brand       string `json:"brand"`
		Status      int    `json:"status"`
	} `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerId        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	Shardkey          string    `json:"shardkey"`
	SmId              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

// here is locate our cache
var dataCache = make(map[string]JsStruct, 10)

func GetDataByID(c *gin.Context) {
	id := c.Param("id")
	res, ok := dataCache[id]
	if ok {
		JSres, err := json.MarshalIndent(res, "", "    ")
		if err != nil {
			panic(err)
		}
		c.HTML(http.StatusOK, "data.html", gin.H{
			"data": string(JSres),
		})
		fmt.Println(err)
		return
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "not found"})
}
func main() {
	jsonData := `{"order_uid": "b563feb7b2b84b6test", "track_number": "WBILMTESTTRACK", "entry": "WBIL", "delivery": { "name": "Muhammed Muhtarov", "phone": "+9720000000", "zip": "2639809", "city": "Kiryat Mozkin", "address": "Ploshad Mira 15", "region": "Kraiot", "email": "test@gmail.com" }, "payment": { "transaction": "b563feb7b2b84b6test", "request_id": "", "currency": "USD", "provider": "wbpay", "amount": 1817, "payment_dt": 1637907727, "bank": "alpha", "delivery_cost": 1500, "goods_total": 317, "custom_fee": 0 }, "items": [ { "chrt_id": 9934930, "track_number": "WBILMTESTTRACK", "price": 453, "rid": "ab4219087a764ae0btest", "name": "Mascaras", "sale": 30, "size": "0", "total_price": 317, "nm_id": 2389212, "brand": "Vivienne Sabo", "status": 202 } ], "locale": "en", "internal_signature": "", "customer_id": "test", "delivery_service": "meest", "shardkey": "9", "sm_id": 99, "date_created": "2021-11-26T06:22:19Z", "oof_shard": "1" }`
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	nc, err := nats.Connect(stan.DefaultNatsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	rows, err := db.Query("SELECT * FROM datawbintern")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	defer rows.Close()

	// Taking data from db and input to the cache
	for rows.Next() {
		var curId string
		var curJSON string
		var curData = JsStruct{}
		err := rows.Scan(&curId, &curJSON)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal([]byte(curJSON), &curData)
		if err != nil {
			panic(err)
		}
		dataCache[curId] = curData
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	fmt.Println("Start in port :8080")
	fmt.Printf("DB Data fetch attempt: %v\n", dataCache)

	sc, err := stan.Connect("test-cluster", "client", stan.NatsURL(stan.DefaultNatsURL))
	if err != nil {
		panic(err)
	}

	err = sc.Publish("publisher", []byte(jsonData))
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		sub, err := sc.Subscribe("publisher", func(m *stan.Msg) {
			curData := JsStruct{}

			fmt.Println("heello world")

			err = json.Unmarshal(m.Data, &curData)
			curId := curData.OrderUid
			dataCache[curId] = curData

			if err != nil {
				return
			}
			_, err = db.Exec("insert into datawbintern (id, data) values ($1, $2)", curId, m.Data)
			if err != nil {
				return
			}
		}, stan.DeliverAllAvailable())
		if err != nil {
			panic(err)
		}
		err = sub.Unsubscribe()
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
	// Close connection
	sc.Close()

	time.Sleep(1 * time.Second)
	r := gin.Default()
	r.LoadHTMLGlob("templates/data.html")

	r.GET("/data/:id", GetDataByID)

	r.Run("localhost:8080")
}
