package main

import (
	"gintest/usfunc"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title string
	Desc string
	Content string
}

func main() {

	r := gin.Default() //定义引擎

	//自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"Add" : usfunc.Add, //自定义函数add -->admin/goods
	})

	//配置html模板文件位置
	r.LoadHTMLGlob("templates/**/*") //templates二层下

	//配置静态文件  第一个表示路由 , 第二个表示对应路径
	r.Static("/kofcss" ,"./kofstatic/static/css")
	r.Static("/kofimage" ,"./kofstatic/static/images")
	r.Static("/kofjs" ,"./kofstatic/js")

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
		c.HTML(http.StatusOK , "default/news.html" , gin.H{
			"title" : "新闻页面",
			"news" : news,
		})
	})
	
	r.GET("/goods" , func(c *gin.Context) {
		c.HTML(http.StatusOK , "admin/goods.html" , gin.H{
			"price" : usfunc.Add(50 , 60),
			"title" : "商品",
		})
	})

	r.GET("/compare" , func(c *gin.Context) {
		c.HTML(http.StatusOK , "default/compare.html" , gin.H{
			"score" : 85,
			"data" : []int{1 , 3 , 5},
		})
	})

	r.GET("/game" , func(c *gin.Context) {
		c.HTML(http.StatusOK , "default/game.html" , gin.H{

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