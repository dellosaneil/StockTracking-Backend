package ws_indicators

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dellosaneil/stocktracking-backend/indicators"
	"github.com/dellosaneil/stocktracking-backend/websockets/ws_price"
	"github.com/gorilla/websocket"
)

var wsVolumeWeightedAveragePriceUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func WSVolumeWeightedAveragePriceUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	wsVolumeWeightedAveragePriceUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := wsVolumeWeightedAveragePriceUpgrader.Upgrade(w, r, nil)

	if err != nil {
		return ws, err
	}
	return ws, nil
}

func WSVolumeWeightedAveragePriceWriter(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Println(t)
			if len(ws_price.WSStockPrice) == 0 {
				continue
			}
			rsi := indicators.VolumeWeightedAveragePrice(ws_price.WSStockPrice)
			jsonString, _ := json.Marshal(rsi)
			if err1 := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err1 != nil {
				panic(err1)
				return
			}
		}
	}
}
