package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers() *httprouter.Router  {
	router := httprouter.New()

	router.POST("/user", CreateUser)

	// 这个router在parse url的时候会把：后面的东西解析为ByName
	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":8000", r)
}

