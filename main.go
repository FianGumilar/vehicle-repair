package main

import (
	"github.com/FianGumilar/vehicle-repair/internal/component"
	"github.com/FianGumilar/vehicle-repair/internal/config"
)

func main() {
	conf := config.Get()

	dbConnection := component.GetDatabaseConnection(conf)
}
