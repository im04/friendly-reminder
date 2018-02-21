package common

import (
	"github.com/kataras/iris"
	"friendly-reminder/manager"
)

// SessShiftExpiration 延时session过期
func SessShiftExpiration(ctx iris.Context) {
	manager.Sess.ShiftExpiration(ctx)
	ctx.Next()
}