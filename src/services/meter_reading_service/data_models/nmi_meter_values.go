package data_models

import "time"

type NMIMeterValues struct {
	NMI         string
	Date        time.Time
	Consumption float64
}

type NMIHeader200Values struct {
	NMI            string
	IntervalLength int8
}
