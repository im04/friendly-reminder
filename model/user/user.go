package user

import (
	"fmt"
	"regexp"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

func CheckName(name string) error {
	flag,err := regexp.MatchString(`^\w{1,8}$`,name)
	if err != nil {
		fmt.Println(err,flag)
		return err
	}
	if !flag {
		return errors.New("用户名不合法")
	}
	return nil
}

func CheckPhone(phone string) error {
	flag,err := regexp.MatchString(`^\d{11}$`,phone)
	if err != nil {
		fmt.Println(err,flag)
		return err
	}
	if !flag {
		return errors.New("用户手机不合法")
	}
	return nil
}

func CheckPassword(ps string) error {
	flag,err := regexp.MatchString(`^\w{8,20}$`,ps)
	if err != nil {
		fmt.Println(err,flag)
		return err
	}
	if !flag {
		return errors.New("用户密码不合法")
	}
	return nil
}

func HaxPassword(password string) string {
	hs := md5.New();
	hs.Write([]byte (password))
	return hex.EncodeToString(hs.Sum(nil))
}