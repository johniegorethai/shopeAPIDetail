package cron

import (
	"context"
	"log"

	"shopeAPIDetail/pkg/firebaseclient"

	"github.com/jmoiron/sqlx"
	"shopeAPIDetail/internal/config"
	"shopeAPIDetail/pkg/httpclient"

	shopeData "shopeAPIDetail/internal/data/shope"

	orderPelapakData "shopeAPIDetail/internal/data/orderPelapak"
	orderPelapakService "shopeAPIDetail/internal/service/orderPelapak"
	//authData "centuryOrder/internal/data/auth"
	//gData "centuryOrder/internal/data/gudang"
	//outData "centuryOrder/internal/data/outlet"
	//plData "centuryOrder/internal/data/pembelianlokal"
	//prData "centuryOrder/internal/data/produk"
	//psData "centuryOrder/internal/data/produksupplier"
	//roData "centuryOrder/internal/data/ro"
	//suppData "centuryOrder/internal/data/supplier"
	//udvbData "centuryOrder/internal/data/uploadDownloadVB"
	//roService "centuryOrder/internal/service/ro"
)

// CRON will auto load configuration, do dependency injection and then start the CRON Job
func CRON() error {
	var (
		//fc    *firebaseclient.Client // Firebase initiation
		//		//hc    *httpclient.Client     // HTTP initialization
		//		//k     *kafka.Kafka           // Kafka Producer
		//		//rd    roData.Data            //
		//		//gd    gData.Data             //
		//		//pld   plData.Data            //
		//		//prd   prData.Data            //
		//		//psd   psData.Data            //
		//		//sd    suppData.Data
		//		//outd  outData.Data      //
		//		//authd authData.Data     //
		//		//udvbd udvbData.Data     //
		//		//rs    roService.Service // BridgingProduct domain service layer
		//		//cfg   *config.Config    // Configuration object

		sd    orderPelapakData.Data // BridgingProduct domain data layer
		cfg   *config.Config        // Configuration object
		sp    shopeData.Data
		fc    *firebaseclient.Client // Firebase initiation
		httpc *httpclient.Client     // HTTP Client
		rs    orderPelapakService.Service
	)
	// Config Init
	//err := config.Init()
	//if err != nil {
	//	log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	//}
	//cfg = config.Get()
	//
	//// Firebase Client Init
	//fc, err = firebaseclient.NewClient(cfg)
	//if err != nil {
	//	log.Fatalf("[FIREBASE] Failed to initialize firebase client: %v", err)
	//}
	//
	//// HTTP Client Init
	//hc = httpclient.NewClient()
	//
	//// Kafka Producer Init
	//k, err = kafka.New(cfg.Kafka.Username, cfg.Kafka.Password, "RO", cfg.Kafka.Brokers)
	//if err != nil {
	//	log.Fatalf("[KAFKA] Failed to initialize kafka producer: %v", err)
	//}
	//
	//// Data layer domain Init
	//authd = authData.New(hc, cfg.API.Auth)
	//gd = gData.New(hc, cfg.API.Gudang)
	//pld = plData.New(hc, cfg.API.PembelianLokal)
	//prd = prData.New(hc, cfg.API.Produk)
	//psd = psData.New(hc, cfg.API.ProdukSupplier)
	//outd = outData.New(hc, cfg.API.Outlet)
	//sd = suppData.New(hc, cfg.API.Supplier)
	//udvbd = udvbData.New(hc, cfg.API.UploadDownloadVB)
	//rd = roData.New(fc)
	//
	//// Service layer domain Init
	//rs = roService.New(rd, k, pld, prd, psd, outd, gd, sd, udvbd, authd)
	//
	//if err := rs.ProsesMalamRO(context.Background()); err != nil {
	//	return err
	//}

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
