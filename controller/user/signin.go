package user

import (
	"github.com/kataras/iris"
	."friendly-reminder/model/user"
	"friendly-reminder/common"
)

func Signin(ctx iris.Context) {
	var signinReq SigninReq
	sendError := common.SendError;
	if err := ctx.ReadJSON(&signinReq); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	//检测用户密码格式
	if  err := CheckSignin(signinReq); err != nil{
		sendError(ctx,err)
		return
	}
	//用户登陆
	data, err := ModelSignin(signinReq)
	if err !=nil {
		sendError(ctx,err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "登陆成功",
		"data": data,
	})
}