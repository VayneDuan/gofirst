package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Golang")
}

//* 两种方式
/*
	- go mod init
	- go mod tidy
	- go build
	* 在当前目录生成bin
*/
/*
	- go mod init
	- go mod tidy
	- go install
	* 在 ~/go/bin生成bin
*/
