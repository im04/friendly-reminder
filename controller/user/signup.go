package user

import (
	"github.com/kataras/iris"
	."friendly-reminder/model/user"
	"friendly-reminder/common"
)

func Signup(ctx iris.Context) {
	var signupReq SignupReq
	sendError := common.SendError;
	if err := ctx.ReadJSON(&signupReq); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	//检查用户资料格式
	if err := CheckSignup(&signupReq); err !=nil {
		sendError(ctx,err)
		return
	}
	//插入用户
	if err := InserUser(&signupReq); err != nil {
		sendError(ctx,err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "注册成功",
	})
}