package header_event_record_parser

import (
	"MeterBilling/src/services/meter_reading_service/data_models"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

const NMI = "NMI2023430"

func getHeaderEvent200ExampleRecord() []string {
	return []string{"200", NMI}
}

func getHeaderEvent300ExampleRecord() []string {
	return []string{"300", "20050301"}
}

func TestHeaderEventRecordParser(t *testing.T) {
	t.Run("header event 200 parser - get nmi", func(t *testing.T) {
		headerEvent200Record := NewHeaderEvent200RecordParser(getHeaderEvent200ExampleRecord())
		require.Equal(t, headerEvent200Record.GetNMI(), NMI)
	})
	t.Run("header event 300 parser - get nmi", func(t *testing.T) {
		headerEvent300Record := NewHeaderEvent300RecordParser(getHeaderEvent300ExampleRecord(), data_models.NMIHeader200Values{NMI: NMI, IntervalLength: 30})
		fmt.Println(headerEvent300Record.GetIntervalDate())
	})
}
