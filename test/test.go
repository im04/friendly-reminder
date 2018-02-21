package test

import (
	"github.com/gorilla/websocket"
)
var list map[string]*websocket.Conn = make(map[string]*websocket.Conn);
func GetList () map[string]*websocket.Conn {
	return list
}