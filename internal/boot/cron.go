package boot

import (
	"context"
	"log"

	"shopeAPIDetail/pkg/firebaseclient"

	"github.com/jmoiron/sqlx"
	"shopeAPIDetail/internal/config"
	"shopeAPIDetail/pkg/httpclient"

	shopeData "shopeAPIDetail/internal/data/shope"

	orderPelapakData "shopeAPIDetail/internal/data/mutasiShopee"
	orderPelapakService "shopeAPIDetail/internal/service/mutasiShopee"
)

// CRON will auto load configuration, do dependency injection and then start the CRON Job
func CRON() error {
	var (
		sd    orderPelapakData.Data // BridgingProduct domain data layer
		cfg   *config.Config        // Configuration object
		sp    shopeData.Data
		fc    *firebaseclient.Client // Firebase initiation
		httpc *httpclient.Client     // HTTP Client
		rs    orderPelapakService.Service
	)

	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg = config.Get()

	fc, err = firebaseclient.NewClient(cfg)
	if err != nil {
		return err
	}

	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	httpc = httpclient.NewClient()

	sp = shopeData.New(httpc, cfg.API.Shope, cfg.API.Shope1, cfg.API.Shope2)
	// BridgingProduct domain init
	sd = orderPelapakData.New(db, fc)

	// Service layer domain Init
	rs = orderPelapakService.New(sd, sp)

	if _, err := rs.InsertMutasiShopeeScheduled(context.Background()); err != nil {
		return err
	}

	return nil
}
