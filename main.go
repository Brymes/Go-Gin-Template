package main

import (
	"App-Name/api"
	"App-Name/config"
)

func init() {
	config.InitDb()
}

func main() {
	api.Server()
}
