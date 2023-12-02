package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// V3

//func Handler(w http.ResponseWriter, r *http.Request) {
//	server := New()
//
//	server.get("/", func(context *Context) {
//		context.JSON(200, H{
//			"message": "test",
//		})
//	})
//
//	server.Handle(w, r)
//}

// V2

var (
	app *gin.Engine
)

type Detail struct {
	T1 string `json:"t1"`
	T2 string `json:"t2"`
}

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"test": "test",
	})
}

func myRoute(r *gin.RouterGroup) {
	r.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("detail", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"T1": "test T1 Lorem Ipsum is simply dummy text of the printing and typesetting",
			"T2": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		})
	})

	r.GET("array", func(c *gin.Context) {
		test := make([]string, 0)
		test = append(test, "test1")
		test = append(test, "test2")

		s := make([]Detail, 0)
		s = append(s, Detail{
			T1: "test T1 Lorem Ipsum is",
			T2: "It is a long established",
		})
		s = append(s, Detail{
			T1: "dummy text of the printing and typesetting",
			T2: "fact that a reader will be distracted by the readable content of a page when looking at its layout.",
		})
		c.JSON(200, gin.H{
			"array":  test,
			"struct": s,
		})
	})

	r.GET("test", test)
}

func init() {
	app = gin.New()
	r := app.Group("/api")
	myRoute(r)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	// Allow requests from any origin

	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Allow specified HTTP methods

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")

	// Allow specified headers

	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept")
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//
	//w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// return "OKOK"
	app.ServeHTTP(w, r)
}

//func enableCORS(next http.Handler) http.Handler {
//
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		// Allow requests from any origin
//
//		w.Header().Set(“Access-Control-Allow-Origin”, “*”)
//
//		// Allow specified HTTP methods
//
//		w.Header().Set(“Access-Control-Allow-Methods”, “GET, POST, OPTIONS”)
//
//		// Allow specified headers
//
//		w.Header().Set(“Access-Control-Allow-Headers”, “Origin, Content-Type, Accept”)
//
//		// Continue with the next handler
//
//		next.ServeHTTP(w, r)
//
//	})
//
//}
//
//// Handler for your API request
//
//func handleAPIRequest(w http.ResponseWriter, r *http.Request) {
//
//	// Handle the API request
//
//	// …
//
//}

// V1

// func main() {
// 	r := gin.Default()
// 	r.GET("/ping", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "pong",
// 		})
// 	})

// 	v1 := r.Group("/v1")
// 	{
// 		v1.GET("/login", loginEndpoint)
// 		v1.GET("/detail", GetDetail)
// 		//v1.POST("/read", readEndpoint)
// 	}

// 	r.Run() // listen and serve on 0.0.0.0:8080
// }

// type StructA struct {
// 	FieldA string `form:"field_a"`
// 	Text   string `form:"text"`
// }

// type Detail struct {
// 	T1 string `json:"t1"`
// 	T2 string `json:"t2"`
// }

// func loginEndpoint(c *gin.Context) {

// 	var b StructA
// 	c.Bind(&b)
// 	c.JSON(200, gin.H{
// 		"field":     b.FieldA,
// 		"Test_text": b.Text,
// 	})
// }

// func GetDetail(c *gin.Context) {
// 	//var d Detail
// 	c.JSON(200, gin.H{
// 		"T1": "test T1 Lorem Ipsum is simply dummy text of the printing and typesetting",
// 		"T2": "It is a long established fact that a reader will be distracted by the readable content of a page when looking at its layout.",
// 	})
// }

// vercel.json
// "routes": [
//   {"src": "/(.*)","dest":"/api"}
// ],
// "build": {
//   "env": {
//     "GO_BUILD_FLAGS": "-ldflags '-s -w'"
//   }
// },
