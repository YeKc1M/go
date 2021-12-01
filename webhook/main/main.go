package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"webhook/model"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "hi")
		return
	})
	r.POST("/payload", func(c *gin.Context) {
		m := make(map[string]interface{})
		c.BindJSON(&m)
		log.Println(m)
		c.String(200, "done")
		return

	})
	r.POST("/json/inline", func(c *gin.Context) {
		var app model.App
		c.Bind(&app)
		log.Println(app)
		c.String(200, "")
		return
	})
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
