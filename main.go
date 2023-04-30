package main

import (
	"App-Name/api"
	"App-Name/config"
)

func init() {
	config.InitDB()
}

func main() {
	api.Server()
}
