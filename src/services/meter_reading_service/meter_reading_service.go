package meter_reading_service

import (
	"MeterBilling/src/logger"
	"MeterBilling/src/services/meter_reading_service/business"
	"MeterBilling/src/services/meter_reading_service/business/header_event_record_executor"
	header_event_record_parser2 "MeterBilling/src/services/meter_reading_service/business/header_event_record_parser"
	meter_reading_service_constants "MeterBilling/src/services/meter_reading_service/constants"
	"MeterBilling/src/services/meter_reading_service/data_models"
	"context"
	"io"
	"os"
)

func StartMeterReadingService(_ctx context.Context) {
	meterReaderService := NewMeterReaderService()
	meterReaderService.Execute()
	return
}
func StopMeterReadingService(_ctx context.Context) {
	// this is empty as it's a simple service
	// As our service will start to process multiple files at scale, this becomes crucial for graceful shutdown
	return
}

type MeterReaderService interface {
	Execute()
}

type meterReaderService struct {
}

func NewMeterReaderService() MeterReaderService {
	return &meterReaderService{}
}

func (m *meterReaderService) Execute() {
	wd, _ := os.Getwd()
	processor := business.NewMeterRecordFileParser(wd + meter_reading_service_constants.METER_RECORD_FILE_READING_PATH)
	var header200Values data_models.NMIHeader200Values

	for {
		nextLine, err := processor.GetNextLine()
		if err == io.EOF {
			logger.GetLogger().Info("finished processing file")
			break
		}
		if err != nil {
			logger.GetLogger().Error("error while processing csv file")
			panic(err)
		}
		recordIndicator := nextLine[0:1][0]
		switch recordIndicator {
		case meter_reading_service_constants.HEADER_RECORD_100:
			continue
		case meter_reading_service_constants.HEADER_RECORD_200:
			Executor200Resp := header_event_record_executor.NewHeaderEvent200Executor(
				header_event_record_parser2.NewHeaderEvent200RecordParser(nextLine),
			).Execute()
			// store the recent visited header 200 values in this struct
			// this will be accessed and passed to header 300 processing
			header200Values = Executor200Resp.(data_models.NMIHeader200Values)
		case meter_reading_service_constants.HEADER_RECORD_300:
			header_event_record_executor.NewHeaderEvent300Executor(
				header_event_record_parser2.NewHeaderEvent300RecordParser(nextLine, header200Values),
				header200Values,
			).Execute()
		case meter_reading_service_constants.HEADER_RECORD_900:
			continue
		}
	}
}
