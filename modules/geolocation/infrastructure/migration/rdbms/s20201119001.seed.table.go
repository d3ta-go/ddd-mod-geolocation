package rdbms

import (
	"fmt"

	domEntity "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/entity"
	"github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/infrastructure/migration/rdbms/data"
	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"gorm.io/gorm"
)

// Seed20201119001InitTable type
type Seed20201119001InitTable struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewSeed20201119001InitTable constructor
func NewSeed20201119001InitTable(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Seed20201119001InitTable)
	gmr.SetHandler(h)
	gmr.SetID("Seed20201119001InitTable")
	return gmr, nil
}

// GetID get Seed20201119001InitTable ID
func (dmr *Seed20201119001InitTable) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Seed20201119001InitTable
func (dmr *Seed20201119001InitTable) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr._seeds(); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Seed20201119001InitTable
func (dmr *Seed20201119001InitTable) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr._unSeeds(); err != nil {
			return err
		}
	}
	return nil
}

func (dmr *Seed20201119001InitTable) _seeds() error {
	if dmr.GetGorm().Migrator().HasTable(&domEntity.CountryEntity{}) {

		if err := dmr.GetGorm().Create(&data.MDM20201119001Countries).Error; err != nil {
			return err
		}

	}
	return nil
}

func (dmr *Seed20201119001InitTable) _unSeeds() error {
	if dmr.GetGorm().Migrator().HasTable(&domEntity.CountryEntity{}) {

		for _, v := range data.MDM20201119001Countries {
			if err := dmr.GetGorm().Unscoped().Where(&v).Delete(&domEntity.CountryEntity{}).Error; err != nil {
				return err
			}
		}

	}
	return nil
}
