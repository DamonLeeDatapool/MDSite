package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"title": "hello gin " + strings.ToLower(c.Request.Method) + " method",
	})
}
