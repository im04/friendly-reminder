package user

import (
	"friendly-reminder/manager"
	"fmt"
)

type SearchUserReq struct {
	Keyword string `json:"keyword"`
}

type SearchUserRep struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
}

func SearchUserModel(keyword string) ([]SearchUserRep,error) {
	db, err := manager.Open()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(keyword)
	rows,err := db.Query("SELECT user_id,user_name FROM USER WHERE user_name LIKE ?","%"+keyword+"%")
	if err != nil {
		fmt.Println("sql")
		fmt.Println(err)
		return nil,err
	}
	var searchUserRep SearchUserRep
	var serachList []SearchUserRep
	for rows.Next() {
		if err := rows.Scan(&searchUserRep.UserId,&searchUserRep.UserName,); err != nil {
			return nil, err
		}
		serachList = append(serachList, searchUserRep)
	}
	return serachList,nil
}
