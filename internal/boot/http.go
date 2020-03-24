package boot

import (
	"log"
	"net/http"
	"shopeAPIDetail/pkg/firebaseclient"

	"github.com/jmoiron/sqlx"
	"shopeAPIDetail/internal/config"
	"shopeAPIDetail/pkg/httpclient"

	shopeData "shopeAPIDetail/internal/data/shope"

	orderPelapakData "shopeAPIDetail/internal/data/orderPelapak"
	orderPelapakServer "shopeAPIDetail/internal/delivery/http"
	orderPelapakHandler "shopeAPIDetail/internal/delivery/http/orderPelapak"
	orderPelapakService "shopeAPIDetail/internal/service/orderPelapak"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {
	var (
		s     orderPelapakServer.Server    // HTTP Server Object
		sd    orderPelapakData.Data        // BridgingProduct domain data layer
		ss    orderPelapakService.Service  // BridgingProduct domain service layer
		sh    *orderPelapakHandler.Handler // BridgingProduct domain handler
		cfg   *config.Config               // Configuration object
		sp    shopeData.Data
		fc  *firebaseclient.Client // Firebase initiation
		httpc *httpclient.Client // HTTP Client
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

	sp = shopeData.New(httpc, cfg.API.Shope,cfg.API.Shope1, cfg.API.Shope2)
	// BridgingProduct domain init
	sd = orderPelapakData.New(db,fc)
	ss = orderPelapakService.New(sd, sp)
	sh = orderPelapakHandler.New(ss)

	// Inject service used on handler
	s = orderPelapakServer.Server{
		OrderPelapak: sh,
	}

	if err := s.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
