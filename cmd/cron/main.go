package cron

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"shopeAPIDetail/internal/cron"
)

func main() {
	if err := cron.CRON; err != nil {
		log.Println("[CRON] failed to cron job due to " + err.Error())
	}
}
