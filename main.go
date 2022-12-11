package main

import (
	"github.com/vijeyash1/backend/pkg/db"
	"github.com/vijeyash1/backend/pkg/routes"
)

func main() {
	db.Init()
	routes.Start()
	
}
