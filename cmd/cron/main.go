package cron

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"shopeAPIDetail/internal/boot"
)

func main() {
	if err := boot.CRON; err != nil {
		fmt.Println("[CRON] failed to cron job due to " + err.Error())
	}
}
