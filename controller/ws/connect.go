package ws

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
	"fmt"
	"friendly-reminder/model/ws"
	"sync"
)

var online = ws.GetOnLineUser()
var mutex = new(sync.Mutex)

func SetupWebsocket(app *iris.Application) {
	// create our echo websocket server
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)
	app.Get("/connect/{userId}", ws.Handler())
	app.Get("/ws", func(ctx iris.Context) {
		ctx.ServeFile("websockets.html", false) // second parameter: enable gzip?
	})
	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

func handleConnection(c websocket.Connection) {
	userId, err := c.Context().Params().GetInt("userId")
	if err != nil {
		c.EmitMessage([]byte(err.Error()))
	}
	fmt.Println(userId)
	online[userId] = ws.User {
		userId,
		c,
	}
	c.OnMessage(func (data []byte) {
		message, err := deCodeJson(data)
		if err != nil {
			fmt.Println(err)
		}
		switch message.EventType {
		case ws.Ev.HeartBeat:
			c.EmitMessage(data)
		}
		fmt.Println(message);
	})
	c.OnDisconnect(func () {
		mutex.Lock()
		delete(online, userId)
		mutex.Unlock()
	})
}
