package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang-seed/internal/handlers"
)

func main() {
	DB := bd.Init()
	
	h := handlers.New(DB)
	
	log.Print("API is running!!")
	http.ListenAndServe(":4000", router)

	r := gin.Default()
	r.GET("/ping", rest.)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}