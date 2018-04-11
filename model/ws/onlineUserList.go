package ws

import (
	"errors"
	"encoding/json"
)

var onLineUser map[int]User

func GetOnLineUser () map[int]User {
	if onLineUser == nil {
		onLineUser = make(map[int]User)
	}
	return onLineUser;
}

func getUserById (userId int) (*User, error){
	if user, ok := onLineUser[userId]; ok {
		return &user, nil
	}
	return nil, errors.New("用户不在线")
}

func SendMsg(fromId int, toId int, msg Message) error {
		fromUser,_ := getUserById(fromId)
		toUser,userErr :=  getUserById(toId)
		str, jsonErr := json.Marshal(msg)
		switch true {
		case userErr != nil:
			return userErr
		case jsonErr != nil:
			return jsonErr
		}
		return sendMsgToUser(fromUser, toUser, str)
}

func sendMsgToUser (fromUser *User, toUser *User, msg []byte) error {
	err := fromUser.Conn.To(toUser.Conn.ID()).EmitMessage(msg)
	if err != nil {
		return err
	}
	return nil
}