package business

import (
	"MeterBilling/src/logger"
	"encoding/csv"
	"os"
	"sync"
)

type MeterRecordFileParser interface {
	GetNextLine() ([]string, error)
}

type meterRecordFileParser struct {
	filePath string
	parser   *csv.Reader
}

var meterRecordFileInstance MeterRecordFileParser
var meterRecordFileInstanceSingleton sync.Once

func NewMeterRecordFileParser(filePath string) MeterRecordFileParser {
	meterRecordFileInstanceSingleton.Do(func() {
		file, err := os.Open(filePath)
		if err != nil {
			logger.GetLogger().Fatal("error while opening meter reading file")
			panic(err)
		}
		parser := csv.NewReader(file)
		// this is necessary to avoid "wrong number of fields" error,
		// our csv has inconsistent number of fields in each row
		parser.FieldsPerRecord = -1
		meterRecordFileInstance = &meterRecordFileParser{
			filePath: filePath,
			parser:   parser,
		}
	})
	return meterRecordFileInstance
}

func (p *meterRecordFileParser) GetNextLine() ([]string, error) {
	record, err := p.parser.Read()
	if err != nil {
		logger.GetLogger().Error("error while parsing csv file")
		return nil, err
	}
	return record, nil
}
