package websockets

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/dellosaneil/stocktracking-backend/api/api_get"
	"github.com/gorilla/websocket"
)

var livePriceUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func LivePriceUpgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	livePriceUpgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := livePriceUpgrader.Upgrade(w, r, nil)
	if err != nil {
		return ws, err
	}
	return ws, nil
}

func LivePriceWriter(conn *websocket.Conn, stockTicker string, timeseries string) {
	for {
		ticker := time.NewTicker(5 * time.Second)
		for t := range ticker.C {
			fmt.Println(t)
			prices, err := api_get.GetPriceCall(stockTicker, timeseries)
			if err != nil {
				panic(err)
			}
			jsonString, _ := json.Marshal(prices)
			if err1 := conn.WriteMessage(websocket.TextMessage, []byte(jsonString)); err != nil {
				panic(err1)
				return
			}
		}
	}
}
