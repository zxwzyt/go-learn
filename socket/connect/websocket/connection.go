package main

import "github.com/gorilla/websocket"

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}
