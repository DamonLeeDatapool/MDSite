package handler

import (
	"net/http"

	"MServer/config"

	"github.com/gin-gonic/gin"
)

type LoginData struct {
	UserName string `form:"username" json:"username" `
	Password string `form:"password" json:"password" `
}

func Login(c *gin.Context) {

	//var dep Department
	//err := c.Bind(&dep)
	//ioutil.ReadAll(c.Request.Body)
	//_trigger := c.Param("TriggerID")
	//a := c.PostForm("TriggerID")

	//name := c.Request.FormValue("name")
	//name := c.Request.PostFormValue("name")
	//因為axios defalut setting content-type=application/json 所以用以下這個 binding:"required"

	var loginData LoginData
	if c.BindJSON(&loginData) == nil {
		config.Logger.Infoln("user:", loginData.UserName)
		if loginData.UserName == "admin" {

			data := struct {
				Id            string   `json:"id"`
				Name          string   `json:"name"`
				Username      string   `json:"username"`
				Password      string   `json:"psaaword"`
				Avatar        string   `json:"avatar"`
				Status        int      `json:"status"`
				Telephone     string   `json:"telephone"`
				LastLoginIp   string   `json:"lastLoingIp"`
				LastLoginTime int64    `json:"lastLoginTime"`
				CreatorId     string   `json:"creatorId"`
				CreateTime    int64    `json:"createTime"`
				MerchantCode  string   `json:"merchantCode"`
				Deleted       int      `json:"deleted"`
				Permission    []string `json:"permission"`
				Token         string   `json:"token"`
				Menu          []string `json:"menu"` //userNav,
			}{
				Id:            "4291d7da9005377ec9aec4a71ea837f",
				Name:          "admin",
				Username:      "Ones@github",
				Password:      "",
				Avatar:        "/avatar2.jpg",
				Status:        1,
				Telephone:     "",
				LastLoginIp:   "27.154.74.117",
				LastLoginTime: 1534837621348,
				CreatorId:     "admin",
				CreateTime:    1497160610259,
				MerchantCode:  "TLif2btpzg079h15bk",
				Deleted:       0,
				Permission:    []string{"admin"},
				Token:         "12312312",
				Menu:          []string{}, //userNav,
			}

			c.JSON(200, data)

		} else {
			config.Logger.Infoln("wrong id and password")
			c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		}
	}

}
