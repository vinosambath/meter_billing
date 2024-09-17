package db

import (
	"sync"

	"gorm.io/gorm"
)

type IGenericORMWrapper interface {
	Insert(data interface{}) error
	FindByNMI(data *IGenericORMWrapper, nmi string) error
}

type genericORMWrapperImpl struct {
	db *gorm.DB
}

var ormInstance IGenericORMWrapper
var ormInstanceSingleton sync.Once

func GetGenericORMWrapper(db *gorm.DB) IGenericORMWrapper {
	ormInstanceSingleton.Do(func() {
		ormInstance = &genericORMWrapperImpl{
			db: db,
		}
	})

	return ormInstance
}

func (d *genericORMWrapperImpl) Insert(data interface{}) error {
	return d.db.Create(data).Error
}

func (d *genericORMWrapperImpl) FindByNMI(data *IGenericORMWrapper, nmi string) error {
	return d.db.Where("nmi = ?", nmi).Find(&data).Error
}
