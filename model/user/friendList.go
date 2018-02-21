package user

import (
	"friendly-reminder/manager"
	"fmt"
)

type FriendListReq struct {
	UserId int `json:"userId"`
}

type friendListRep struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPhone int `json:"userPhone"`
	IsAgree int `json:"-"`
}

func GetFriendList(userId int) ([]friendListRep,[]friendListRep, error) {
	fmt.Print(userId)
	db, err := manager.Open()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	// 获取好友列表
	rows, err := db.Query("SELECT user.user_id,user_name,user_phone,is_agree FROM user,user_user WHERE user_user.user_id = ? AND user.user_id = user_user.friend_id AND user_user.is_agree = 1",userId)
	defer rows.Close()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	var friendList []friendListRep
	var fr friendListRep
	for rows.Next() {
		if err = rows.Scan(&fr.UserId,&fr.UserName,&fr.UserPhone,&fr.IsAgree);err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		friendList = append(friendList,fr)
	}

	//获取好友请求
	rs, err := db.Query("SELECT user.user_id,user_name,user_phone,is_agree FROM user,user_user WHERE user_user.friend_id = ? AND user.user_id = user_user.user_id AND user_user.is_agree = 0", userId)
	defer rs.Close()
	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}
	var questList []friendListRep
	var qu friendListRep
	for rs.Next() {
		if err = rs.Scan(&qu.UserId,&qu.UserName,&qu.UserPhone,&qu.IsAgree);err != nil {
			fmt.Println(err)
			return nil, nil, err
		}
		questList = append(questList,qu)
	}
	fmt.Println(questList, friendList);
	return friendList, questList, nil
}
