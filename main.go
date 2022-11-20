package main

import (
	"build/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func main() {
	r := gin.Default()
	r.Use(middleware.CORSMiddleware())
	r.LoadHTMLGlob("./tmpl/*")
	r.GET("/build", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})

	})
	r.POST("/build", func(c *gin.Context) {
		u := c.PostForm("username")
		p := c.PostForm("password")
		fmt.Println(string(u), string(p))

		if u != "admin" || p != "1111" {
			c.JSON(200, gin.H{
				"msg": "404",
			})
		} else {

			c.HTML(http.StatusOK, "build.html", gin.H{
				"username": u,
			})
		}

	})

	r.GET("/dobuild", func(c *gin.Context) {
		//cmd := exec.Command("cmd", "/C", "dir")
		//
		//cmd.Stdout = os.Stdout
		//cmd.Stderr = os.Stderr
		//
		//e := cmd.Run()
		//if e != nil {
		//	fmt.Println(e)
		//}
		out, _ := exec.Command("cmd", "/C", "dir").Output()
		outs := string(out)

		c.JSON(200, gin.H{
			"message": "命令已经执行，请等待输出结果。。。。。。。。",
			"info":    outs,
		})
	})
	r.Run()
}
