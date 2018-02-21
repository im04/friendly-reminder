package user

import (
	"friendly-reminder/common"
	."friendly-reminder/model/user"
	"github.com/kataras/iris"
)


func Accept(ctx iris.Context)  {
	var accept AcceptReq
	sendError := common.SendError;
	if err := ctx.ReadJSON(&accept); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	if err := AcceptApplication(&accept);err != nil {
		sendError(ctx, err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "已添加为好友",
	})
}