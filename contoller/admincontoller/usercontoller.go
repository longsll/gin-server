package admincontoller

import (
	"fmt"
	"gintest/models"
	"net/http"
	_ "os/user"

	"github.com/gin-gonic/gin"
)

type Usercontoller struct {
	Basecontoller
}

func (con Usercontoller) Index(c *gin.Context){

	userlist := []models.User{}
	models.DB.Find(&userlist)
	c.JSON(http.StatusOK , gin.H{
		"result":userlist,
	})

	c.String(200,"index")
}

func (con Usercontoller) Add(c *gin.Context){
	user := models.User{
		User: "test",
		Age: 18,
		Email: "111@qq.com",
		Addtime: 20033,
	}
	models.DB.Create(&user)
	fmt.Println(user)
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
	user.Email = "edit@qq.com"

	models.DB.Save(&user)

	c.String(200,"edit")
}