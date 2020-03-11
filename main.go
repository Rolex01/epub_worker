package main

import (
	"flag"
	"fmt"
	"github.com/rolex01/epub_workers/epub"
	"github.com/rolex01/epub_workers/fb2"
	"log"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8080, "API service port")
}

func main() {
	flag.Parse()

	router := fasthttprouter.New()
	router.POST("/api/get/fb2_parse", fb2.FB2Parse)
	router.POST("/api/get/epub_parse", epub.EpubParse)

	log.Printf("starting server as :%d\n", port)
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf(":%d", port), router.Handler))
}