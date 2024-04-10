package main

import (
	"github.com/eduhub/ws"
	"net/http"
)

func main() {
	http.HandleFunc("/ws", ws.GetWebSocketManager().InitChatRoom)
	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
