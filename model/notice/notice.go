package notice

import (
	"friendly-reminder/manager"
	"fmt"
	"errors"
)

type Notice struct {
	FromId int `json:"fromId"`
	ToId int `json:"toId"`
	NoticeTitle string `json:"noticeTitle"`
	NoticeContent string `json:"noticeContent"`
	NoticeTime int64 `json:"noticeTime"`
}

type GetNoticeListRep struct {
	UserId int `json:"userId"`
	ListType int `json:"listType"`
}

func IntserNotice (rep Notice) error {
	db, err := manager.Open()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	//插入消息
	_,err = db.Exec("INSERT INTO notice (from_user_id,to_user_id,notice_title,notice_content,notice_time)" +
		"VALUE (?,?,?,?,?)",rep.FromId,rep.ToId,rep.NoticeTitle,rep.NoticeContent,rep.NoticeTime)
	if err != nil {
		return err
	}
	return nil
}

func GetNoticeList(userId int, listType int) ([]Notice, error) {
	var lt string
	switch true {
	case listType == 1:
		lt = "from_user_id"
	case listType == 2:
		lt = "to_user_id"
	default:
		return nil, errors.New("listType参数错误")
	}
	db, err := manager.Open()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(lt);
	//获取消息列表
	rows, err := db.Query("SELECT from_user_id,to_user_id,notice_title,notice_content,notice_time FROM notice WHERE notice." + lt + " = ?",userId)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var noticeList []Notice
	var noticeReader Notice
	for rows.Next() {
		if err = rows.Scan(&noticeReader.FromId,&noticeReader.ToId,&noticeReader.NoticeTitle,&noticeReader.NoticeContent,&noticeReader.NoticeTime);err != nil {
			fmt.Println(err)
			return nil, err
		}
		noticeList = append(noticeList, noticeReader)
	}
	fmt.Println(noticeList);
	return noticeList, nil
}