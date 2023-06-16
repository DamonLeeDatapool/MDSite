package models

import (
	"MServer/config"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Pwd        string `json:"pwd"`
	CreateTime string `json:"createtime"`
}

func (u *User) TableName() string {
	return "User"
}

func GetAllUser(user *[]User) (err error) {
	if err = config.DB.Find(user).Error; err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) (err error) {
	if err = config.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserByID(user *User, id string) (err error) {
	if err = config.DB.Where("id = ?", id).First(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *User, id string) (err error) {
	//fmt.Println(user)
	config.DB.Save(user)
	return nil
}

func DeleteUser(user *User, id string) (err error) {
	config.DB.Where("id = ?", id).Delete(user)
	return nil
}
