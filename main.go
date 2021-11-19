package main

import (
	"Golang/Golang_Post/myapp"
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("프로그램 실행")

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
