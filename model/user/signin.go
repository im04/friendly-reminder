package user

import (
	"friendly-reminder/manager"
	"fmt"
	"errors"
)

type SigninReq struct {
	UserName string `json:"userName"`
	UserPassword string `json:"userPassword"`
}

type SigninQuery struct {
	userId int `json:"userId"`
	userName string `json:"userName"`
	userPhone int `json:"userPhone"`
	userPassword string `json:"userPhone"`
}

type signinRep struct {
	UserId int `json:"userId"`
	UserName string `json:"userName"`
	UserPhone int `json:"userPhone"`
}

func ModelSignin(user SigninReq) (*signinRep,error) {
	db, err := manager.Open()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	row,err := db.Query("SELECT * FROM USER WHERE user_name=?",user.UserName)
	defer func () {
		if err == nil {
			row.Close()
		}
	}()
	if err != nil {
		return nil, errors.New("用户查询错误")
	}
	if !row.Next(){
		return nil, errors.New("用户不存在")
	}
	var userData SigninQuery
	err = row.Scan(&userData.userId,&userData.userName,&userData.userPhone,&userData.userPassword)
	if err != nil {
		return nil, errors.New("用户查询错误")
	}
	if (HaxPassword(user.UserPassword) == userData.userPassword) {
		data := signinRep {
			UserId: userData.userId,
			UserName: userData.userName,
			UserPhone: userData.userPhone,
		}
		return &data, nil
	} else {
		return nil, errors.New("用户密码错误")
	}
	return nil, errors.New("用户查询错误")
}

func CheckSignin(user SigninReq) error {
	errU := CheckName(user.UserName)
	errPw := CheckPassword(user.UserPassword)
	switch true {
	case errU != nil:
		return errU
	case errPw != nil:
		return errPw
	default:
		return nil
	}
}