package main

import (
	"github.com/eduhub/configs"
	"github.com/eduhub/ws"
	"net/http"
)

func main() {
	Db, err := configs.SetupDB()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/ws", ws.GetWebSocketManager(Db).InitChatRoom)
	err = http.ListenAndServe(":8889", nil)
	if err != nil {
		panic("Error starting server: " + err.Error())
	}
}
