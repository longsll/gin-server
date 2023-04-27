package route

import (
	"gintest/contoller/defaultcontoller"
	"gintest/contoller/admincontoller"
	"github.com/gin-gonic/gin"
)

type Article struct {
	Title string
	Desc string
	Content string
}

func Defaultinit(r *gin.Engine) {
	defaultrouter := r.Group("/")
	{
		//中间件
		defaultrouter.GET("/" ,defaultcontoller.Defcon{}.Initmiddleware, defaultcontoller.Defcon{}.Deffront)

		defaultrouter.GET("/news" , defaultcontoller.Defcon{}.Defnews)
		
		defaultrouter.GET("/goods" , defaultcontoller.Defcon{}.Defgoods)

		defaultrouter.GET("/compare" , defaultcontoller.Defcon{}.Defcompare)	
	}
}

func Usertinit(r *gin.Engine) {
	userrouter := r.Group("/user")
	{
		userrouter.GET("/index" , admincontoller.Usercontoller{}.Index)	
		userrouter.GET("/add" , admincontoller.Usercontoller{}.Add)	
		userrouter.GET("/delete" , admincontoller.Usercontoller{}.Delete)	
		userrouter.GET("/edit" , admincontoller.Usercontoller{}.Edit)	
	}
}