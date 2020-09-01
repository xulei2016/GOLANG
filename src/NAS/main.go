package main

import (
	"HTTP"
)

func main() {
	HTTP.StartLog()

	//开启web监听
	HTTP.StartSever()
}