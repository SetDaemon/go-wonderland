package main

import (
	"github.com/setdaemon/go-wonderland/sdweb/spider"
	"log"
	"net/http"
)

func main() {
	log.Println("main run")
	s := spider.New()
	s.GET("/", func(c *spider.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Neo<h1>")
	})
	s.GET("/hello", func(c *spider.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	s.POST("/login", func(c *spider.Context) {
		c.JSON(http.StatusOK, spider.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	s.Run(":9999")
}

//func helloHandler(w http.ResponseWriter, req *http.Request) {
//	for k, v := range req.Header {
//		fmt.Fprintf(w, "111Header[%q] = %q\n", k, v)
//	}
//}
//
//func indexHandler(res http.ResponseWriter, req *http.Request) {
//	fmt.Fprintf(res, "111URL.Path = %q\n", req.URL)
//}
//
//type Engine struct{}
//
//func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	switch req.URL.Path {
//	case "/":
//		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
//	case "/hello":
//		for k, v := range req.Header {
//			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
//		}
//	default:
//		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
//	}
//}
