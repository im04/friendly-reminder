package user

import (
	"friendly-reminder/manager"
	"fmt"
	"errors"
)

type AcceptReq struct {
	UserId int `json:"userId"`
	FriendId int `json:"friendId"`
}

func AcceptApplication(user *AcceptReq) error {
	db, err := manager.Open()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	res,err := db.Exec("UPDATE user_user SET is_agree=1 WHERE user_id=? AND friend_id=? AND is_agree=0",user.FriendId,user.UserId)
	if err != nil {
		fmt.Println(err)
		return err
	}
	cl,err := res.RowsAffected()
	if err !=nil {
		fmt.Println(err)
		return err
	}
	if cl == 0 {
		fmt.Println(cl)
		return errors.New("该用户已经是好友")
	}
	friendRequestReq := FriendRequestReq{
		user.UserId,
		user.FriendId,
	}
	if err := FriendRequestModel(friendRequestReq);err != nil {
		return err
	}
	res,err = db.Exec("UPDATE user_user SET is_agree=1 WHERE user_id=? AND friend_id=? AND is_agree=0",user.UserId,user.FriendId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}