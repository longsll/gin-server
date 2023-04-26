package defaultcontoller

import (
	"fmt"
	"gintest/usfunc"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title string
	Desc string
	Content string
}

type Defcon struct{}

//输出执行时间啊
func (con Defcon) Initmiddleware(c *gin.Context) {
	start := time.Now().UnixNano()

	c.Next() //相当于对下面的执行defer

	end := time.Now().UnixNano()

	fmt.Println(end - start)
}

func (con Defcon) Deffront(c *gin.Context) {
	c.HTML(http.StatusOK , "default/index.html" , gin.H{
		"title" : "首页",
	})
}

func (con Defcon) Defnews(c *gin.Context) {
	news := &Article{
		Title: "新闻标题",
		Desc: "简述",
		Content: "内容",
	}
	c.HTML(http.StatusOK , "default/news.html" , gin.H{
		"title" : "新闻页面",
		"news" : news,
	})
}

func (con Defcon) Defgoods(c *gin.Context) {
	c.HTML(http.StatusOK , "admin/goods.html" , gin.H{
		"price" : usfunc.Add(50 , 60),
		"title" : "商品",
	})
}

func (con Defcon) Defcompare(c *gin.Context) {
	c.HTML(http.StatusOK , "default/compare.html" , gin.H{
		"score" : 85,
		"data" : []int{1 , 3 , 5},
	})
}

