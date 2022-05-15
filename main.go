package main

import (
	"github.com/Zaida-3dO/goblin/adapters/driver/rest/routes"
	"github.com/Zaida-3dO/goblin/config"
)

func main() {
	config.Init()
	routes.Init()
}
