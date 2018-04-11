package user

import (
	"fmt"
	"friendly-reminder/manager"
	"errors"
)

type FriendRequestReq struct {
	UserId int `json:"userId"`
	FriendId int `json:"friendId"`
}

func FriendRequestModel (user FriendRequestReq) error {
	db, err := manager.Open()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if user.UserId <= 0 || user.FriendId <= 0 {
		return errors.New("用户Id和好友Id不能为空")
	}

	if user.UserId == user.FriendId {
		return errors.New("不能添加自己为好友")
	}

	//_,err = db.Exec("INSERT INTO user_user (user_id,friend_id)" + "VALUE (?,?)",user.UserId,user.FriendId)

	res,err := db.Exec("INSERT INTO user_user (user_id,friend_id) SELECT ?,? FROM dual WHERE not exists (select * from user_user where user_id=? and friend_id=?)",
		user.UserId,user.FriendId,
		user.UserId,user.FriendId);
	if err != nil {
		return err
	}

	cl,err := res.RowsAffected()
	if err !=nil {
		fmt.Println(err)
		return err
	}
	if cl == 0 {
		fmt.Println(cl)
		return errors.New("不能重复发送好友申请")
	}

	return nil
}