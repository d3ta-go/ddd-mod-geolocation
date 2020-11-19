package rdbms

import (
	"fmt"

	domEntity "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/entity"
	"github.com/d3ta-go/system/system/handler"
	migRDBMS "github.com/d3ta-go/system/system/migration/rdbms"
	"gorm.io/gorm"
)

// Migrate20201119001InitTable type
type Migrate20201119001InitTable struct {
	migRDBMS.BaseGormMigratorRunner
}

// NewMigrate20201119001InitTable constructor
func NewMigrate20201119001InitTable(h *handler.Handler) (migRDBMS.IGormMigratorRunner, error) {
	gmr := new(Migrate20201119001InitTable)
	gmr.SetHandler(h)
	gmr.SetID("Migrate20201119001InitTable")
	return gmr, nil
}

// GetID get Migrate20201119001InitTable ID
func (dmr *Migrate20201119001InitTable) GetID() string {
	return fmt.Sprintf("%T", dmr)
}

// Run run Migrate20201119001InitTable
func (dmr *Migrate20201119001InitTable) Run(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if err := dmr.GetGorm().AutoMigrate(
			&domEntity.CountryEntity{},
		); err != nil {
			return err
		}
	}
	return nil
}

// RollBack rollback Migrate20201119001InitTable
func (dmr *Migrate20201119001InitTable) RollBack(h *handler.Handler, dbGorm *gorm.DB) error {
	if dbGorm != nil {
		dmr.SetGorm(dbGorm)
	}
	if dmr.GetGorm() != nil {
		if dmr.GetGorm().Migrator().HasTable(&domEntity.CountryEntity{}) {
			if err := dmr.GetGorm().Migrator().DropTable(&domEntity.CountryEntity{}); err != nil {
				return err
			}
		}
	}
	return nil
}
