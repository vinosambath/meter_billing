package header_event_record_parser

import (
	"MeterBilling/src/logger"
	"strconv"
)

type HeaderEvent200RecordParser interface {
	GetNMI() string
	GetIntervalLength() int8
}

type headerEvent200RecordParser struct {
	record []string
}

func NewHeaderEvent200RecordParser(record []string) HeaderEvent200RecordParser {
	return &headerEvent200RecordParser{
		record: record,
	}
}

func (h *headerEvent200RecordParser) GetNMI() string {
	if len(h.record) < 2 {
		logger.GetLogger().Fatal("incorrect csv file format - 200 row - get nmi")
		panic("invalid file with 200 header row")
	}
	return h.record[1:2][0]
}

func (h *headerEvent200RecordParser) GetIntervalLength() int8 {
	if len(h.record) < 8 {
		logger.GetLogger().Fatal("incorrect csv file format - 200 row - incorrect interval length")
		panic("invalid file with 200 header row")
	}
	intervalLength, err := strconv.Atoi(h.record[8:9][0])
	if err != nil {
		panic("invalid interval length")
	}
	return int8(intervalLength)
}
