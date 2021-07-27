package main

import (
	"ken.com.tw/demo/internal/router"
)

func main() {
	r := router.Initial()
	r.Run(":8888")
}