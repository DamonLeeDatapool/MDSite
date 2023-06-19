package models

import (
	"MServer/config"
	//"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id           uint   `json:"id" gorm:"cloumn:User_Id; primaryKey"`
	UserAccount  string `json:"user_account" gorm:"column:User_Account"`
	userPassword string `json:"user_password" gorm:"column:User_Password"`
	userName     string `json:"user_name gorm:"column:User_Name"`
	userPhone    string `json:"user_phone gorm:"column:User_Phone"`
	StoreId      int    `json:"store_id gorm:"column:Store_Id"`
	CreateTime   int64  `json:"createtime" gorm:"column:CreateTime"`
	UpdateTime   int64  `json:"updatetime" gorm:"column:UpdateTime"`
}

func (u *User) TableName() string {
	return "users"
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
