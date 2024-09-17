package main

import (
	"MeterBilling/src/configuration"
	"MeterBilling/src/db"
	"MeterBilling/src/logger"
	"MeterBilling/src/services/meter_reading_service"

	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	configuration.InitConfig()
	logger.InitLogger()
	db.InitDBConnections()

	_, cancel := start()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	// allow graceful shutdown of independent services
	defer shutdown(cancel)

}

func start() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithCancel(context.Background())
	meter_reading_service.StartMeterReadingService(ctx)
	return
}

func shutdown(cancel context.CancelFunc) {
	cancel()
	ctx, cancelTimeout := context.WithTimeout(context.Background(), time.Second*30)
	meter_reading_service.StopMeterReadingService(ctx)
	defer cancelTimeout()
}
