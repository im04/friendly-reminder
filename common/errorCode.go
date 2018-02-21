package common

import "github.com/kataras/iris"

var ErrorCode map[string]iris.Map = map[string]iris.Map {
	//参数错误
	"paramsFail": iris.Map{
		"errCode": "1001",
		"errType": "参数错误",
	},
	//注册
	"signupFail": iris.Map{
		"errCode": "1011",
		"errType": "注册失败",
	},
	//登陆
	"signinFail": iris.Map{
		"errCode": "1021",
		"errType": "登陆失败",
	},
	//好友列表
	"friendListFail": iris.Map{
		"errCode": "1031",
		"errType": "获取好友列表失败",
	},
}

