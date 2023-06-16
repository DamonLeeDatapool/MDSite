package handler

import (
	"net/http"
	"path/filepath"
	"strconv"
	"time"

	"MServer/config"
	"MServer/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Maintain_Item
func GetMaintainItemList(c *gin.Context) {
	var moItem []models.MaintainItem
	err := models.GetAllMaintainItem(&moItem)
	if err != nil {
		config.Logger.Println("get maintain_item error:", err)
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, moItem)
	}
}

func GetMaintainItemListByFT(c *gin.Context) {
	var moItem []models.MaintainItem
	ft := c.Params.ByName("ft")
	err := models.GetMaintainItemByFixerType(&moItem, ft)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, moItem)
	}
}

// Maintain_Order
func GetMaintainOrderAll(c *gin.Context) {
	var mo []models.MaintainOrder
	err := models.GetAllMaintainOrderWithLimit(&mo, 50) //limit=50 ,暫時先寫
	if err != nil {
		config.Logger.Println("get maintain_order error:", err)
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, mo)

}

func CreateMaintainOrder(c *gin.Context) {
	var mo models.MaintainOrder
	mo.MaintainItemId, _ = strconv.Atoi(c.PostForm("maintain_item_id"))
	mo.UserId, _ = strconv.Atoi(c.PostForm("user_id"))
	mo.AppointDate, _ = time.Parse("2006-01-02", c.PostForm("appoint_date"))
	tempValue, _ := strconv.Atoi(c.PostForm("appoint_time"))
	mo.AppointTime = int8(tempValue)
	mo.Location = c.PostForm("location")
	mo.Comment = c.PostForm("Comment")
	mo.Status = 0
	mo.CreateTime = time.Now().Unix()
	mo.UpdateTime = mo.CreateTime

	id, err := models.CreateMaintainOrder(&mo)
	if err != nil {
		config.Logger.Println("create maintain_order error:", err)
		c.AbortWithStatus(http.StatusNotFound)
	}

	//存file
	var moFile models.MaintainOrderFile
	file, _ := c.FormFile("file")
	ext := filepath.Ext(file.Filename)
	path := "static/pic/" + strconv.Itoa(id) + "/"
	fileName := uuid.New().String() + ext
	uploadUrl := path + fileName
	if err := c.SaveUploadedFile(file, uploadUrl); err != nil {
		config.Logger.Printf("SaveUploadedFile,err=%v", err)
		c.String(http.StatusBadRequest, "上傳圖片失敗")
		return
	}

	moFile.MaintainOrderId = id
	moFile.FileName = fileName
	moFile.FilePath = uploadUrl
	moFile.FileContentType = ext
	moFile.CreateTime = mo.CreateTime
	moFile.UpdateTime = mo.CreateTime

	err = models.CreateMaintainOrderFile(&moFile)

	if err != nil {
		config.Logger.Println("create maintain_order error:", err)
		c.AbortWithStatus(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, "新增維修成功")

}
