package main

import (
	"github.com/pieterclaerhout/go-log"
	"github.com/voltento/mongo_rest/app/config"
	"github.com/voltento/mongo_rest/app/metrics"
	"github.com/voltento/mongo_rest/app/repo"
	"github.com/voltento/mongo_rest/app/service"
	"github.com/voltento/mongo_rest/app/web"
)

func main() {
	log.DebugMode = true
	log.DebugSQLMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	hc := service.NewHealthChecker()

	cfg := config.GetConfig()

	r := repo.NewMongo(cfg)
	defer r.Disconnect()

	search := service.NewSearch(r)
	hc.RegisterService(r)

	s := web.NewServer(cfg, search, hc)

	p := metrics.NewPrometheus(s.Router())
	p.Use(s.Router())

	err := s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
