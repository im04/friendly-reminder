package ws

import (
	"github.com/kataras/iris"
	"fmt"
	"encoding/json"
	."friendly-reminder/model/ws"
	"friendly-reminder/common"
	"friendly-reminder/model/notice"
	"github.com/kataras/iris/core/errors"
)

func PushMsg(ctx iris.Context) {
	var rep notice.Notice
	sendError := common.SendError;
	if err := ctx.ReadJSON(&rep); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	paramsNull := ""
	switch true {
	case rep.FromId == 0:
		paramsNull = "fromId不能为空";
	case rep.ToId == 0:
		paramsNull = "toId不能为空";
	case rep.NoticeTitle == "":
		paramsNull = "noticeTitle不能为空";
	case rep.NoticeTitle == "":
		paramsNull = "noticeTitle不能为空";
	case rep.NoticeTime == 0:
		paramsNull = "noticeTime不能为空";
	}
	if paramsNull != "" {
		sendError(ctx, errors.New(paramsNull),"paramsFail")
		return
	}
	err := SendMsg(rep.FromId, rep.ToId, Message{
		Ev.SetTips,
		rep,
	})
	if err !=nil {
		sendError(ctx,err)
		return
	}
	err = notice.IntserNotice(rep)
	if err !=nil {
		sendError(ctx,err)
		return
	}
	fmt.Println(rep)
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "发送消息成功",
	})
}

func enCodeJson(jsonData interface{}) (*[]byte, error) {
	str, err := json.Marshal(jsonData)
	if err != nil {
		return nil, err
	}
	return &str, nil
}

func deCodeJson(s []byte) (*Message, error) {
	msg := Message{}
	err := json.Unmarshal(s, &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}