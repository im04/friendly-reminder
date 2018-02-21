package user

import (
	"github.com/kataras/iris"
	"friendly-reminder/common"
	."friendly-reminder/model/user"
)

func FriendRequest(ctx iris.Context) {
	var friendRequestReq FriendRequestReq
	sendError := common.SendError;
	if err := ctx.ReadJSON(&friendRequestReq); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	if err := FriendRequestModel(friendRequestReq);err != nil {
		sendError(ctx, err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "申请好友成功",
	})
}