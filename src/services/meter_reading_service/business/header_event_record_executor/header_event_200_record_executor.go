package header_event_record_executor

import (
	"MeterBilling/src/services/meter_reading_service/business/header_event_record_parser"
	"MeterBilling/src/services/meter_reading_service/data_models"
)

type headerEvent200RecordExecutor struct {
	headerEvent200RecordParser header_event_record_parser.HeaderEvent200RecordParser
}

func NewHeaderEvent200Executor(headerEvent200RecordParser header_event_record_parser.HeaderEvent200RecordParser) HeaderEventRecordExecutor {
	return &headerEvent200RecordExecutor{
		headerEvent200RecordParser: headerEvent200RecordParser,
	}
}

func (h *headerEvent200RecordExecutor) Execute() interface{} {
	nmi := h.headerEvent200RecordParser.GetNMI()
	intervalLength := h.headerEvent200RecordParser.GetIntervalLength()

	return data_models.NMIHeader200Values{
		NMI:            nmi,
		IntervalLength: intervalLength,
	}
}
