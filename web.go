package main

import (
	"fmt"
	"net/http"

	"github.com/dchest/captcha"
	"github.com/go-martini/martini"
	"github.com/mojo-zd/vim-tips-web/routers"
)

var (
	m = martini.Classic()
)

func main() {
	routers.Initialize(m)

	http.HandleFunc("/ws", routers.WSHandler)
	http.Handle("/captcha", captcha.Server(captcha.StdWidth, captcha.StdHeight))
	http.Handle("/", m)

	fmt.Println("Server started...")

	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
