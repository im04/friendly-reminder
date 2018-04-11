package route

import (
	"github.com/kataras/iris"
	"friendly-reminder/controller/user"
	"friendly-reminder/controller/ws"
	"friendly-reminder/controller/notice"
)



func Route(app *iris.Application) {
	ws.SetupWebsocket(app)
	routes := app.Party("/api")
	{
		routes.Post("/signin", user.Signin)
		routes.Post("/signup", user.Signup)
		routes.Post("/friendList", user.FriendList)
		routes.Post("/accept",user.Accept)
		routes.Post("/searchUser",user.SearchUser)
		routes.Post("/friendRequest",user.FriendRequest)
		routes.Post("/pushNotice", ws.PushMsg)
		routes.Post("/noticeList", notice.NoticeList)
	}
}