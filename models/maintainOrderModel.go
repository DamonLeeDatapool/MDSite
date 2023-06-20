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

type MaintainOrderInfoReturnToFront struct {
	MaintainOrderId   int    `json:"maintainOrderId"`
	MaintainItemName  string `json:"maintainItemName"`
	FixerType         string `json:"fixerType"`
	OrderStatus       int    `json:"orderStatus"`
	OrderStatusString string `json:"orderStatusString"`
	CreateTime        int64  `json:"createTime"`
	CreateTimeString  string `json:"createTimeString"`
	TimeElpased       int    `json:"timeElasped` //CreateTime後經過幾小時
	OvertimeStatus    string `json:"overtimeStatus"`
}

type MaintainOrderDetailInfoReturnToFront struct {
	MaintainOrderId   int       `json:"maintainOrderId"`
	MaintainItemName  string    `json:"maintainItemName"`
	UserName          string    `json:"userName"`
	UserPhone         string    `json:"userPhone"`
	AppointDate       time.Time `json:"appointDate"`
	AppointTime       int       `json:"appointTime"`
	AppointTimeString string    `json:"appointTimeString"`
	Comment           string    `json:"comment"`
	//OrderStatusNo    int    `json:"orderStatusNo"`
	//OrderStatus      string `json:"orderStatus"`
	PicUrl string `json:"picUrl`
	//CreateTime       int    `json:"createTime"`
	//TimeElpased      int    `json:"timeElasped` //CreateTime後經過幾小時
	//OvertimeStatue   int    `json:"overtimeStatus"`
}

var AppointTimeMapping = map[int]string{
	0: "早上",
	1: "下午",
	2: "晚上",
}

var OrderStatusMapping = map[int]string{
	0: "預約中",
	1: "處理中",
	2: "已解決",
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

func GetAllMaintainOrderWithLimitForFront(MOExtItem *[]MaintainOrderInfoReturnToFront, limit int) (err error) {

	fields := `
	maintain_order.Maintain_Order_Id AS maintain_order_id,
	maintain_item.Name AS maintain_item_name, 
	maintain_item.Fixer_Type AS fixer_type, 
	maintain_order.CreateTime AS create_time, 
	maintain_order.Status AS order_status 	
	`

	if err = config.DB.Table("maintain_order").Select(fields).
		Joins("left join maintain_item on maintain_order.Maintain_Item_Id = maintain_item.Id").
		Order("create_time desc").Limit(limit).Scan(MOExtItem).Error; err != nil {
		return err
	}

	MaintainOrderPostProcess(MOExtItem)

	return nil
}

func GetMaintainOrderStatusForFront(MOExtItem *[]MaintainOrderInfoReturnToFront, status int) (err error) {

	fields := `
	maintain_order.Maintain_Order_Id AS maintain_order_id,
	maintain_item.Name AS maintain_item_name, 
	maintain_item.Fixer_Type AS fixer_type, 
	maintain_order.CreateTime AS create_time, 
	maintain_order.Status AS order_status 	
	`

	if err = config.DB.Table("maintain_order").Select(fields).
		Joins("left join maintain_item on maintain_order.Maintain_Item_Id = maintain_item.Id").
		Where("Status = ?", status).
		Order("create_time desc").Scan(MOExtItem).Error; err != nil {
		return err
	}

	MaintainOrderPostProcess(MOExtItem)

	return nil
}

func MaintainOrderPostProcess(MOExtItem *[]MaintainOrderInfoReturnToFront) {

	for i := 0; i < len(*MOExtItem); i++ {
		//OrdetStatusString
		(*MOExtItem)[i].OrderStatusString = OrderStatusMapping[(*MOExtItem)[i].OrderStatus]
		//OverTimeStattus
		TimeElpasedSeconds := time.Now().Unix() - (*MOExtItem)[i].CreateTime
		(*MOExtItem)[i].TimeElpased = int(TimeElpasedSeconds / 3600)
		if (*MOExtItem)[i].TimeElpased > 48 && (*MOExtItem)[i].OrderStatus == 0 {
			(*MOExtItem)[i].OvertimeStatus = "已超時"
		} else if (*MOExtItem)[i].TimeElpased > 24 && (*MOExtItem)[i].OrderStatus == 0 {
			(*MOExtItem)[i].OvertimeStatus = "即將超時"
		} else {
			(*MOExtItem)[i].OvertimeStatus = "符合規定"
		}
		//CreateTimeString
		(*MOExtItem)[i].CreateTimeString = time.Unix((*MOExtItem)[i].CreateTime, 0).Format("2006-01-02 15:04:05")
	}

}

func GetOneMaintainOrderWithLimitForFront(MOExtItem *MaintainOrderDetailInfoReturnToFront, id int) (err error) {

	fields := `
	maintain_order.Maintain_Order_Id AS maintain_order_id, 
	maintain_item.Name AS maintain_item_name, 
	users.User_Name AS user_name,
	users.User_Phone AS user_phone,
	maintain_order.Appoint_Date AS appoint_date,
	maintain_order.Appoint_Time AS appoint_time,
	maintain_order.Comment AS comment,
	maintain_order_file.FilePath AS pic_url	
	`

	if err = config.DB.Table("maintain_order").Select(fields).
		Joins("left join maintain_item on maintain_order.Maintain_Item_Id = maintain_item.Id").
		Joins("left join maintain_order_file on maintain_order.Maintain_Order_Id = maintain_order_file.Maintain_Order_Id").
		Joins("left join users on maintain_order.User_Id = users.User_Id").
		Where("maintain_order.Maintain_Order_Id = ? ", id).Scan(MOExtItem).Error; err != nil {
		return err
	}

	MOExtItem.AppointTimeString = AppointTimeMapping[MOExtItem.AppointTime]
	return nil
}

func CreateMaintainOrder(mo *MaintainOrder) (id int, err error) {
	if err = config.DB.Create(mo).Error; err != nil {
		return 0, err
	}

	return mo.MaintainOrderId, nil
}

func DeleteMaintainorder(mo *MaintainOrder, id int) (err error) {
	config.DB.Where("Maintain_Order_Id = ?", id).Delete(mo)
	return nil
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
