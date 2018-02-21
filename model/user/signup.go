package user

import (
	"database/sql"
	"errors"
	"friendly-reminder/manager"
	"log"
)

type SignupReq struct {
	UserName string `json:"userName"`
	UserPassword string `json:"userPassword"`
	UserPhone string `json:"userPhone"`
}

func InserUser(user *SignupReq) error{
	//打开数据库链接
	db, err := manager.Open()
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}

	//账号是否已经存在
	if err := hasUserName(db, user.UserName); err != nil {
		return err
	}
	//插入用户
	_,err = db.Exec("INSERT INTO user (user_name,user_phone,user_password)" +
		"VALUE (?,?,?)",user.UserName,user.UserPhone,HaxPassword(user.UserPassword))
	if err != nil {
		return err
	}
	return nil
}

func CheckSignup(user *SignupReq) error {
	errU := CheckName(user.UserName)
	errP := CheckPhone(user.UserPhone)
	errPw := CheckPassword(user.UserPassword)
	switch true {
	case errU != nil:
		return errU
	case errP != nil:
		return errP
	case errPw != nil:
		return errPw
	default:
		return nil
	}
}
func hasUserName(db * sql.DB, user_name string) error {
	rows,err := db.Query("SELECT user_name FROM USER WHERE user_name=?",user_name)
	defer rows.Close()
	if err != nil {
		return err
	}
	if rows.Next() {
		return errors.New("用户名已经存在")
	}
	return nil
}