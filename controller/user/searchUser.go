package user

import (
	"github.com/kataras/iris"
	."friendly-reminder/model/user"
	"friendly-reminder/common"
	"strings"
)

func SearchUser(ctx iris.Context) {
	var searchUserReq SearchUserReq
	sendError := common.SendError;
	if err := ctx.ReadJSON(&searchUserReq); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	list,err := SearchUserModel(strings.TrimSpace(searchUserReq.Keyword))
	if err != nil {
		sendError(ctx,err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "查询用户成功",
		"data": list,
	})
}
