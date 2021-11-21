package main

import (
	"github.com/gin-gonic/gin"
	"log"
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
	err := r.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
