package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/login", loginEndpoint)
		v1.GET("/detail", GetDetail)
		//v1.POST("/read", readEndpoint)
	}

	r.Run() // listen and serve on 0.0.0.0:8080
}

type StructA struct {
	FieldA string `form:"field_a"`
	Text   string `form:"text"`
}

type Detail struct {
	T1 string `json:"t1"`
	T2 string `json:"t2"`
}

func loginEndpoint(c *gin.Context) {

	var b StructA
	c.Bind(&b)
	c.JSON(200, gin.H{
		"field":     b.FieldA,
		"Test_text": b.Text,
	})
}

func GetDetail(c *gin.Context) {
	//var d Detail
	c.JSON(200, gin.H{
		"T1": "test T1 Lorem Ipsum is simply dummy text of the printing and typesetting",
		"T2": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
	})
}
