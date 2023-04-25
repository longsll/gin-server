package route

import (
	"gintest/usfunc"
	"net/http"

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
		defaultrouter.GET("/" , func(c *gin.Context) {
			c.HTML(http.StatusOK , "default/index.html" , gin.H{
				"title" : "首页",
			})
		})

		defaultrouter.GET("/news" , func(c *gin.Context) {
			news := &Article{
				Title: "新闻标题",
				Desc: "简述",
				Content: "内容",
			}
			c.HTML(http.StatusOK , "default/news.html" , gin.H{
				"title" : "新闻页面",
				"news" : news,
			})
		})
		
		defaultrouter.GET("/goods" , func(c *gin.Context) {
			c.HTML(http.StatusOK , "admin/goods.html" , gin.H{
				"price" : usfunc.Add(50 , 60),
				"title" : "商品",
			})
		})

		defaultrouter.GET("/compare" , func(c *gin.Context) {
			c.HTML(http.StatusOK , "default/compare.html" , gin.H{
				"score" : 85,
				"data" : []int{1 , 3 , 5},
			})
		})	
	}
}