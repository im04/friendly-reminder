package ws

import (
	"github.com/kataras/iris/websocket"
	"friendly-reminder/model/notice"
)

const (
	heartBeat string = "HEART_BEAT"
	setTips string = "SET_TIPS"
)

type ev struct {
	HeartBeat string
	SetTips string
}

var Ev = ev {
	heartBeat,
	setTips,
}

type User struct {
	UserId int
	Conn websocket.Connection
}
type Message struct {
	EventType string `json:"eventType"`
	Data notice.Notice `json:"data"`
}

