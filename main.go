package main

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

func main() {

	r := gin.Default() //定义引擎
	//配置html模板文件位置
	r.LoadHTMLGlob("templates/**/*") //templates二层下
	//返回字符串
	// r.GET("/", func(c *gin.Context) {
	// 	c.String(200,"zhi","hello world")
	// })
	//返回json
	r.GET("/json1" ,func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//加载html前端
	r.GET("/news" , func(c *gin.Context) {
		news := &Article{
			Title: "新闻标题",
			Desc: "简述",
			Content: "内容",
		}
		c.HTML(http.StatusOK , "admin/news.html" , gin.H{
			"title" : "新闻页面",
			"news" : news,
		})
	})
	
	r.GET("/goods" , func(c *gin.Context) {
		c.HTML(http.StatusOK , "admin/goods.html" , gin.H{
			"price" : 20,
			"title" : "商品",
		})
	})

	r.GET("/compare" , func(c *gin.Context) {
		c.HTML(http.StatusOK , "default/compare.html" , gin.H{
			"score" : usfunc.Add(50 , 60),
			"data" : []int{1 , 3 , 5},
		})
	})

	r.GET("/" , func(c *gin.Context) {
		c.HTML(http.StatusOK , "default/index.html" , gin.H{
			"title" : "首页",
		})
	})
	r.Run(":8080") // 监听并在 0.0.0.0:8080 上启动服务
	//go get github.com/pilu/fresh 热加载
}