package ws_indicators

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dellosaneil/stocktracking-backend/indicators"
	"github.com/dellosaneil/stocktracking-backend/util"
	"github.com/dellosaneil/stocktracking-backend/websockets/ws_price"
	"github.com/gorilla/websocket"
)

var wsExponentialMovingAverageUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSExponentialMovingAverageUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	wsExponentialMovingAverageUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := wsExponentialMovingAverageUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return ws, err
	}
	return ws, nil
}

func WSExponentialMovingAverageWriter(conn *websocket.Conn, period int, priceType string) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Println(t)
			prices := util.PriceType(ws_price.WSStockPrice, priceType)
			if len(prices) == 0 {
				continue
			}
			ema := indicators.ExponentialMovingAverage(prices, period)
			jsonString, _ := json.Marshal(ema)
			if err1 := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err1 != nil {
				panic(err1)
				return
			}
		}
	}
}
