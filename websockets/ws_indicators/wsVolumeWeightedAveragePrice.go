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
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func WSVolumeWeightedAveragePriceUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := wsVolumeWeightedAveragePriceUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return ws, err
	}
	return ws, nil
}

func WSVolumeWeightedAveragePriceWriter(conn *websocket.Conn) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for range ticker.C {
			if len(ws_price.WSStockPrice) > 0 {
				vwap := indicators.VolumeWeightedAveragePrice(ws_price.WSStockPrice)
				jsonString, _ := json.Marshal(vwap)
				if err1 := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err1 != nil {
					fmt.Println(err1)
					return
				}
			}
		}
	}
}
