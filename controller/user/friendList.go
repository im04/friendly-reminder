package user

import (
	"github.com/kataras/iris"
	."friendly-reminder/model/user"
	"friendly-reminder/common"
)


func FriendList(ctx iris.Context) {
	var friendListReq FriendListReq
	sendError := common.SendError;
	if err := ctx.ReadJSON(&friendListReq); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	friendData,requestData,err := GetFriendList(friendListReq.UserId)
	if err !=nil {
		sendError(ctx,err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "获取好友列表成功",
		"requestData": requestData,
		"friendData": friendData,
	})
}
