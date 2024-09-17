package header_event_record_executor

import (
	"MeterBilling/src/db/meter_reading_tables"
	"MeterBilling/src/logger"
	"MeterBilling/src/services/meter_reading_service/business/header_event_record_parser"
	"MeterBilling/src/services/meter_reading_service/data_models"

	"github.com/google/uuid"
)

type headerEvent300Executor struct {
	header_parser     header_event_record_parser.HeaderEvent300RecordParser
	header_200_values data_models.NMIHeader200Values
}

func NewHeaderEvent300Executor(
	header_parser header_event_record_parser.HeaderEvent300RecordParser,
	header_200_values data_models.NMIHeader200Values,
) HeaderEventRecordExecutor {
	return &headerEvent300Executor{
		header_parser:     header_parser,
		header_200_values: header_200_values,
	}
}

func (h *headerEvent300Executor) Execute() interface{} {
	nmi := h.header_200_values.NMI
	intervalDate := h.header_parser.GetIntervalDate()
	consumption := h.header_parser.GetConsumption()

	meterReadingRecord := meter_reading_tables.MeterReadingsTable{
		ID:          uuid.New().String(),
		NMI:         nmi,
		Timestamp:   intervalDate,
		Consumption: consumption,
	}
	err := meterReadingRecord.Insert()
	if err != nil {
		logger.GetLogger().Error("error while saving meter reading record")
	}

	return nil
}
