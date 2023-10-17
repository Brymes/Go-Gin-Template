package main

import (
	"App-Name/api"
	"App-Name/config"
)

func init() {
	config.LoadEnv()
	config.InitDB()
}

func main() {
	api.Server()
}
