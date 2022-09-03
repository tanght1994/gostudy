package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 添加全局处理函数(中间件)
	r.Use(aaa, bbb)

	r.GET("/abc", a, b, c)
	r.GET("/def", d, e, f)
	r.GET("/xyz", x, y, z)

	r.GET("/fun1", fun1, fun2)
	r.GET("/fun2", fun2)

	v1 := r.Group("/v1", aaa)
	v1.POST("/fun1", fun1)
	v1.POST("/fun2", fun2)

	v2 := r.Group("/v2", bbb)
	v2.POST("/fun1", fun1)
	v2.POST("/fun2", fun2)

	// 正则路由
	r.GET("/user/:name", regexp_url)

	r.Run("0.0.0.0:8000")
}

func a(c *gin.Context) {
	fmt.Println("a start")
	fmt.Println("a end")
}

func b(c *gin.Context) {
	fmt.Println("b start")
	fmt.Println("b end")
}

func c(c *gin.Context) {
	fmt.Println("c start")
	c.String(200, "OK")
	fmt.Println("c end")
}

func d(c *gin.Context) {
	fmt.Println("d start")
	c.String(200, "OK-d")
	fmt.Println("d end")
}

func e(c *gin.Context) {
	fmt.Println("e start")
	c.String(200, "OK-e")
	fmt.Println("e end")
}

func f(c *gin.Context) {
	fmt.Println("f start")
	c.String(200, "OK-f")
	fmt.Println("f end")
}

func x(c *gin.Context) {
	fmt.Println("x start")
	c.Next()
	fmt.Println("x end")
}

func y(c *gin.Context) {
	fmt.Println("y start")
	c.Next()
	fmt.Println("y end")
}

func z(c *gin.Context) {
	fmt.Println("z start")
	c.String(200, "OK")
	fmt.Println("z end")
}

func aaa(c *gin.Context) {
	fmt.Println("aaa start")
	c.Next()
	fmt.Println("aaa end")
}

func bbb(c *gin.Context) {
	fmt.Println("bbb start")
	c.Next()
	fmt.Println("bbb end")
}

func fun1(c *gin.Context) {
	fmt.Println("fun1 start")
	c.JSON(http.StatusOK, gin.H{"data": "fun1"})
	fmt.Println("fun1 end")
}

func fun2(c *gin.Context) {
	fmt.Println("fun2 start")
	c.Next()
	c.JSON(http.StatusOK, gin.H{"data": "fun2"})
	fmt.Println("fun2 end")
}

func regexp_url(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{"data": name})
}

func arg_in_url(c *gin.Context) {
	// 如果k1不存在则返回空字符串
	v1 := c.Query("k1")
	// 如果k1不存在则返回"default value"
	v2 := c.DefaultQuery("k1", "default value")
	c.JSON(http.StatusOK, gin.H{"data": fmt.Sprintf("v1[%s] v2[%s]", v1, v2)})
}

func arg_in_post(c *gin.Context) {
	v1 := c.PostForm("k1")
	v2 := c.PostFormArray("k1")
	fmt.Println(v1)
	fmt.Println(v2)
	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}
