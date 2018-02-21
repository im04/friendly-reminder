package common

import (
	"github.com/kataras/iris"
)

func SendError(ctx iris.Context, errMsg error, errName ...string) {
	length := len(errName);
	if length > 0 {
		err := ErrorCode[errName[0]]
		ctx.JSON(iris.Map{
			"errCode" : err["errCode"],
			"errType"   : err["errType"],
			"errDetail": errMsg.Error(),
		})
	} else {
		ctx.JSON(iris.Map{
			"errCode" : "1002",
			"errType"   : "程序错误",
			"errDetail": errMsg.Error(),
		})
	}

}
