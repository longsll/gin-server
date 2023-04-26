package route

import (
	"gintest/contoller/defaultcontoller"

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
		//加载html前端

		//中间件
		defaultrouter.GET("/" ,defaultcontoller.Defcon{}.Initmiddleware, defaultcontoller.Defcon{}.Deffront)

		defaultrouter.GET("/news" , defaultcontoller.Defcon{}.Defnews)
		
		defaultrouter.GET("/goods" , defaultcontoller.Defcon{}.Defgoods)

		defaultrouter.GET("/compare" , defaultcontoller.Defcon{}.Defcompare)	
	}
}