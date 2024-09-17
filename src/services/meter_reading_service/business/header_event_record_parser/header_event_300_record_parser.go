package header_event_record_parser

import (
	"MeterBilling/src/logger"
	"MeterBilling/src/services/meter_reading_service/constants"
	"MeterBilling/src/services/meter_reading_service/data_models"
	"strconv"
	"time"
)

type HeaderEvent300RecordParser interface {
	GetIntervalDate() time.Time
	GetConsumption() float64
}

type headerEvent300RecordParser struct {
	record          []string
	header200Values data_models.NMIHeader200Values
}

func NewHeaderEvent300RecordParser(record []string, header200Values data_models.NMIHeader200Values) HeaderEvent300RecordParser {
	return &headerEvent300RecordParser{
		record:          record,
		header200Values: header200Values,
	}
}

func (h *headerEvent300RecordParser) GetIntervalDate() time.Time {
	if len(h.record) < 2 {
		logger.GetLogger().Fatal("incorrect csv file format - 300 row - interval date")
		panic("invalid file with 300 header row")
	}
	timeStamp, _ := time.Parse("20060102", h.record[1:2][0])
	return timeStamp
}

func (h *headerEvent300RecordParser) GetConsumption() float64 {
	var consumption float64

	intervalRecordsFromMeterStarting := h.record[2:]
	intervalRecords := constants.MINUTES_IN_A_DAY / int(h.header200Values.IntervalLength)
	if len(intervalRecordsFromMeterStarting) < int(intervalRecords) {
		logger.GetLogger().Fatal("incorrect csv file format - 300 row - not enough interval records")
		panic("invalid file with 300 header row")
	}

	for _, record := range intervalRecordsFromMeterStarting[:intervalRecords] {
		intervalRecord, err := strconv.ParseFloat(record, 64)
		if err != nil {
			logger.GetLogger().Error("error while processing csv file - invalid interval field value")
			continue
		}
		consumption += intervalRecord
	}

	return consumption
}
