package main

import (
	"fmt"

	"github.com/ducktyst/goapi/internal/config"
	"github.com/ducktyst/goapi/internal/database"
	"github.com/ducktyst/goapi/internal/server"
)

func main() {
	cfg, err := config.ReadConfig("../configs/config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	db, err := database.New(cfg.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	dborm, err := database.NewDBorm(cfg.Database.ConnectionString)
	if err != nil {
		panic(err)
	}
	defer dborm.Close()

	fmt.Println(dborm)
	s := server.New(cfg.Server, db)

	fmt.Println(db)
	s.Run()
}
