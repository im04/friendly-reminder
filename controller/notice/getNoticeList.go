package notice

import (
	"github.com/kataras/iris"
	"friendly-reminder/common"
	"friendly-reminder/model/notice"
	"errors"
)

func NoticeList(ctx iris.Context) {
	var getNoticeListRep notice.GetNoticeListRep
	sendError := common.SendError;
	if err := ctx.ReadJSON(&getNoticeListRep); err != nil {
		sendError(ctx, err,"paramsFail")
		return
	}
	paramsNull := ""
	switch true {
	case getNoticeListRep.UserId == 0:
		paramsNull = "userId不能为空";
	case getNoticeListRep.ListType == 0:
		paramsNull = "listType不能为空";
	}
	if paramsNull != "" {
		sendError(ctx, errors.New(paramsNull),"paramsFail")
		return
	}
	noticeData,err := notice.GetNoticeList(getNoticeListRep.UserId, getNoticeListRep.ListType)
	if err !=nil {
		sendError(ctx,err)
		return
	}
	ctx.JSON(iris.Map{
		"errCode": "1000",
		"msg": "获取消息列表成功",
		"noticeData": noticeData,
	})
}