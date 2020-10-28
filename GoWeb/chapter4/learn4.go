package main

import (
	"fmt"
	"net/http"
)

func learn4() {

	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/set/cookie", setCookie)
	http.HandleFunc("/get/cookie", showCookie)
	server.ListenAndServe()
}

func setCookie(w http.ResponseWriter, r *http.Request) {

	c1 := http.Cookie{
		Name:  "cookie1",
		Value: "wcj",
	}

	http.SetCookie(w, &c1)
}

func showCookie(w http.ResponseWriter, r *http.Request) {

	// c1 := r.Header.Get("Cookie")
	// if c1 == "" {
	// 	fmt.Fprintln(w, "方法一获取cookie信息失败")
	// } else {
	// 	fmt.Fprintln(w, c1)
	// }

	c2, err := r.Cookie("cookie1")
	if err != nil {
		fmt.Fprintln(w, c2)
	} else {
		fmt.Fprintln(w, "方法二获取cookie信息失败")
	}

	c3 := r.Cookies()
	if c3 == nil {
		fmt.Fprintln(w, "方法三获取cookie失败")
	} else {
		fmt.Fprintln(w, c3)
	}
}
