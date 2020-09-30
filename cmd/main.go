package main

import (
	"fmt"
	"github.com/ducktyst/goapi/internal/config"
	"github.com/ducktyst/goapi/internal/database"
	"github.com/ducktyst/goapi/internal/server"
	"github.com/ducktyst/goapi/internal/update_rate"
	"time"
)

func main() {
	cfg, err := config.ReadConfig("/Users/aleksej/Documents/Универ/ЦПИбву-21/coe_internet_programmirovanie/exchange/configs/config.toml")
	if err != nil {
		panic(err)
	}
	fmt.Println(cfg)

	db, err := database.New(cfg.Database)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	go func() {
		for {
			update_rate.UpdateRates(db)
			time.Sleep(time.Hour)
		}
	}()

	s := server.New(cfg.Server, db)
	s.Run()

}
