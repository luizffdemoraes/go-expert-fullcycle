package main

import "github.com/fullcycle/curso-go/7-Apis/configs"

func main() {
	config, _ := configs.LoadConfig(".")
	println(config.DBDriver)
}
