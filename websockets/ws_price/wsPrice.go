package ws_price

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dellosaneil/stocktracking-backend/api/api_get"
	"github.com/dellosaneil/stocktracking-backend/model"
	"github.com/gorilla/websocket"
)

var wsPriceUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var WSStockPrice []model.PriceModel

func WSPriceUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := wsPriceUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return ws, err
	}
	return ws, nil
}

func WSPriceWriter(conn *websocket.Conn, stockTicker string, timeseries string) {
	err := requestPrice(conn, stockTicker, timeseries)
	if err != nil {
		return
	}
	for {
		ticker := time.NewTicker(60 * time.Second)
		for range ticker.C {
			err := requestPrice(conn, stockTicker, timeseries)
			if err != nil {
				return
			}
		}
	}
}

func requestPrice(conn *websocket.Conn, stockTicker string, timeseries string) error {
	fmt.Println("price")
	prices, err := api_get.GetPriceCall(stockTicker, timeseries)
	WSStockPrice = prices
	if err != nil {
		panic(err)
	}
	jsonString, _ := json.Marshal(prices)
	if err1 := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err1 != nil {
		return err1
	}
	return nil
}
