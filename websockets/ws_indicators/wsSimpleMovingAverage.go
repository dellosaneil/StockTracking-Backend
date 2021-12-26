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

var wsSimpleMovingAverageUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WSSimpleMovingAverageUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := wsSimpleMovingAverageUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return ws, err
	}
	return ws, nil
}

func WSSimpleMovingAverageWriter(conn *websocket.Conn, period int, priceType string) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Println(t)
			prices := util.PriceType(ws_price.WSStockPrice, priceType)
			if len(prices) > 0 {
				fmt.Println("shouldnt")
				sma := indicators.SimpleMovingAverage(prices, period)
				jsonString, _ := json.Marshal(sma)
				if err1 := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err1 != nil {
					fmt.Println(err1)
					return
				}
			}
		}
	}
}
