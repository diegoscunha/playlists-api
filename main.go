package main

import (
	"playlits-music/api/configs"
	"playlits-music/api/router"
)

func main() {
	configs.Init()

	router.Init()
}
