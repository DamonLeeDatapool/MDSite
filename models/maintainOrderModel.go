package models

import (
	"time"
	//"fmt"

	"MServer/config"

	_ "github.com/go-sql-driver/mysql"
)

// MaintainItem
type MaintainItem struct {
	Id        int    `json:"id"`
	Category  string `json:"category"`
	Name      string `json:"name"`
	FixerType string `json:"fixer_type" gorm:"column:Fixer_Type"`
}

func (u *MaintainItem) TableName() string {
	return "maintain_item" //real table name in DB
}

func GetAllMaintainItem(MOItem *[]MaintainItem) (err error) {
	if err = config.DB.Find(MOItem).Error; err != nil {
		return err
	}
	return nil
}

func GetMaintainItemByFixerType(MOItem *[]MaintainItem, FT string) (err error) {
	if err = config.DB.Where("fixer_type = ?", FT).Find(MOItem).Error; err != nil {
		return err
	}
	return nil
}

// MaintainOrder
type MaintainOrder struct {
	MaintainOrderId int       `json:"maintain_order_id" gorm:"cloumn:Maintain_Orde_Id; primaryKey"`
	MaintainItemId  int       `json:"maintain_item_id" gorm:"cloumn:Maintain_Item_Id"`
	UserId          int       `json:"user_id" gorm:"column:User_Id"`
	AppointDate     time.Time `json:"appoint_date" gorm:"column:Appoint_Date"`
	AppointTime     int8      `json:"appoint_time" gorm:"column:Appoint_Time"`
	Location        string    `json:"location"`
	Comment         string    `json:"Comment"`
	Status          int8      `json:"status"`
	CreateTime      int64     `json:"createtime" gorm:"column:CreateTime"`
	UpdateTime      int64     `json:"updatetime" gorm:"column:UpdateTime"`
}

func (u *MaintainOrder) TableName() string {
	return "maintain_order"
}

func GetAllMaintainOrderWithLimit(MOItem *[]MaintainOrder, limit int) (err error) {
	if err = config.DB.Order("CreateTime desc").Limit(limit).Find(MOItem).Error; err != nil {
		return err
	}
	return nil
}

func CreateMaintainOrder(mo *MaintainOrder) (id int, err error) {
	if err = config.DB.Create(mo).Error; err != nil {
		return 0, err
	}

	return mo.MaintainOrderId, nil
}

// MaintainOrderFile
type MaintainOrderFile struct {
	Id              int    `json:"id"`
	MaintainOrderId int    `json:"maintain_order_id"`
	FileName        string `json:"filename" gorm:"column:FileName"`
	FilePath        string `json:"filepath" gorm:"column:FilePath"`
	FileContentType string `json:"file_content_type" gorm:"column:FileContentType"`
	CreateTime      int64  `json:"createtime" gorm:"column:CreateTime"`
	UpdateTime      int64  `json:"updatetime" gorm:"column:UpdateTime"`
}

func (u *MaintainOrderFile) TableName() string {
	return "maintain_order_file"
}

func CreateMaintainOrderFile(moFile *MaintainOrderFile) (err error) {
	if err = config.DB.Create(moFile).Error; err != nil {
		return err
	}

	return nil
}
