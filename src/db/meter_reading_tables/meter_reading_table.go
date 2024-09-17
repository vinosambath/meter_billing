package meter_reading_tables

import (
	"MeterBilling/src/configuration"
	"MeterBilling/src/db"
	"MeterBilling/src/utils"
	"time"

	"gorm.io/gorm"
)

type MeterReadingsTable struct {
	ID          string
	NMI         string
	Timestamp   time.Time
	Consumption float64
}

func (t *MeterReadingsTable) Insert() error {
	return db.GetGenericORMWrapper(t.computeShard(t.NMI)).Insert(t)
}

func (t *MeterReadingsTable) FindByNMI(data *db.IGenericORMWrapper, nmi string) error {
	return db.GetGenericORMWrapper(t.computeShard(nmi)).FindByNMI(data, nmi)
}

func (t *MeterReadingsTable) computeShard(nmi string) *gorm.DB {
	config := configuration.GetConfig()
	numberOfDatabases := uint32(config.Database.NumberOfDatabases)
	return db.GetDBConnections()[utils.GenStringToHash(nmi)%numberOfDatabases]
}
