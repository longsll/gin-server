package admincontoller

import (
	"fmt"
	"gintest/models"
	"net/http"
	_ "os/user"

	"github.com/gin-gonic/gin"
	_ "gorm.io/gorm"
)

type Usercontoller struct {
	Basecontoller
}

func (con Usercontoller) Defaultpage(c *gin.Context){
	
	c.HTML(http.StatusOK , "default/user.html" , gin.H{
	})
	//username := c.Query("username")
	// user := models.User{}
	// user.User = username
	// models.DB.Create(&user)
}

func (con Usercontoller) Adduser(c *gin.Context){
	tesuser := c.PostForm("username")
	c.JSON(http.StatusOK , gin.H{
		"username" :tesuser,
	})
}

func (con Usercontoller) Logger(c *gin.Context){
	user := models.User{}
	ans := models.User{}
	user.User = c.PostForm("username")
	user.Pwd = c.PostForm("password")
	models.DB.Where("user = ? AND pwd = ?", user.User, user.Pwd).Find(&ans)
	if user.User != ans.User || user.Pwd != ans.Pwd {
		c.HTML(http.StatusOK , "default/user.html" , gin.H{})
	}else{
		c.HTML(http.StatusOK , "default/aduser.html" , gin.H{})
	}
}

func (con Usercontoller) Index(c *gin.Context){
	userlist := []models.User{}
	models.DB.Find(&userlist)
	c.JSON(http.StatusOK , gin.H{
		"result":userlist,
	})
	c.String(200,"index")
}


func (con Usercontoller) Addef(c *gin.Context){
	c.HTML(http.StatusOK , "default/aduser.html" , gin.H{})
}

func (con Usercontoller) Adddeau(c *gin.Context){
	c.HTML(http.StatusOK , "default/adduser.html" , gin.H{})
}

func (con Usercontoller) Addding(c *gin.Context){

	user := models.User{}
	user.User = c.PostForm("username")
	user.Pwd = c.PostForm("password")
	models.DB.Create(&user)
	fmt.Println(user.User)
	fmt.Println(user.Pwd)
	c.String(200,"add")
	
}



func (con Usercontoller) Delete(c *gin.Context){

	user := models.User{Id: 6}

	models.DB.Delete(&user)

	c.String(200,"delete")
}

func (con Usercontoller) Edit(c *gin.Context){
	
	user := models.User{Id: 5}

	models.DB.Find(&user)
	
	user.User = "testedit1"

	models.DB.Save(&user)

	c.String(200,"edit")
}