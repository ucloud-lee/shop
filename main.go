package main

import (
	"log"
	"net/http"

	"git.ucloudadmin.com/leesin/shop/pkg/account"
	"git.ucloudadmin.com/leesin/shop/pkg/shoes"
	"git.ucloudadmin.com/leesin/shop/pkg/user"
)

func main() {
	http.HandleFunc("/index", user.Index)
	http.HandleFunc("/login/login.html", account.Login)

	// shoes
	http.HandleFunc("/shoes/adidasi.html", shoes.Adidasi)
	http.HandleFunc("/shoes/nike.html", shoes.Nike)
	http.HandleFunc("/shoes/vans.html", shoes.Vans)
	http.HandleFunc("/shoes/airjodan.html", shoes.AirJodan)

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
