package controllers

import (
	"blogweb_gin/models"
	"blogweb_gin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RegisterGet(c *gin.Context) {
	// return html
	c.HTML(http.StatusOK, "register.html", gin.H{"title": "注册页"})
}

//处理注册
func RegisterPost(c *gin.Context) {
	//获取表单信息
	username := c.PostForm("username")
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	fmt.Println(username, password, repassword)

	//注册之前判断该用户名是否已经被注册，如果已经注册，则返回错误
	id := models.QueryUserWithUsername(username)
	fmt.Println("id:", id)
	if id > 0 {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "用户名已存在"})
		return
	}
	//注册用户名和密码
	//登录的时候进行密码md5校验
	password = utils.MD5(password)
	fmt.Println("MD5后的密码为：", password)

	user := models.User{0, username, password, 0, time.Now().Unix()}
	_, err := models.InsertUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 0, "message": "注册失败"})
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "message": "注册成功"})
	}
}
