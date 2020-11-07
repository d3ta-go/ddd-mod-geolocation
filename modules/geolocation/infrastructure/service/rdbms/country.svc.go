package rdbms

import (
	"fmt"
	"net/http"
	"strings"

	domEntity "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/entity"
	domSchema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
	sysErr "github.com/d3ta-go/system/system/error"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
)

// NewCountrySvc new CountrySvc
func NewCountrySvc(h *handler.Handler) (*CountrySvc, error) {
	svc := new(CountrySvc)
	svc.SetHandler(h)

	cfg, err := h.GetDefaultConfig()
	if err != nil {
		return nil, err
	}
	svc.SetDBConnectionName(cfg.Databases.MainDB.ConnectionName)

	return svc, nil
}

// CountrySvc represent CountrySvc
type CountrySvc struct {
	BaseSvc
}

// ListAll list All Country
func (s *CountrySvc) ListAll(i identity.Identity) (*domSchema.ListCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryList []domEntity.CountryEntity
	if err := dbCon.Order("code asc").Find(&countryList).Error; err != nil {
		return nil, err
	}

	// response
	var listCountry []*domSchema.Country
	for _, rec := range countryList {
		tmp := new(domSchema.Country)

		tmp.ID = rec.ID
		tmp.Code = rec.Code
		tmp.Name = rec.Name
		tmp.ISO2Code = rec.ISO2Code
		tmp.ISO3Code = rec.ISO3Code
		tmp.WHORegion = rec.WHORegion

		listCountry = append(listCountry, tmp)
	}

	res := new(domSchema.ListCountryResponse)
	res.Data = listCountry

	return res, nil
}

// GetDetail get Detail Country by Code
func (s *CountrySvc) GetDetail(req *domSchema.GetDetailCountryRequest, i identity.Identity) (*domSchema.GetDetailCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryEtt domEntity.CountryEntity
	if err := dbCon.Where("code = ?", req.Code).First(&countryEtt).Error; err != nil || countryEtt.Code == "" {
		return nil, &sysErr.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Data not found (for Code=%s)", req.Code)}
	}

	// response
	country := &domSchema.Country{
		ID:        countryEtt.ID,
		Code:      countryEtt.Code,
		Name:      countryEtt.Name,
		ISO2Code:  countryEtt.ISO2Code,
		ISO3Code:  countryEtt.ISO3Code,
		WHORegion: countryEtt.WHORegion,
	}

	res := new(domSchema.GetDetailCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}

// Add Country
func (s *CountrySvc) Add(req *domSchema.AddCountryRequest, i identity.Identity) (*domSchema.AddCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// save
	countryEtt := &domEntity.CountryEntity{
		Code:      req.Code,
		Name:      req.Name,
		ISO2Code:  req.ISO2Code,
		ISO3Code:  req.ISO3Code,
		WHORegion: req.WHORegion,
	}
	countryEtt.CreatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Create(countryEtt).Error; err != nil {
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysErr.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, err
	}

	// response
	country := &domSchema.Country{
		ID:        countryEtt.ID,
		Code:      countryEtt.Code,
		Name:      countryEtt.Name,
		ISO2Code:  countryEtt.ISO2Code,
		ISO3Code:  countryEtt.ISO3Code,
		WHORegion: countryEtt.WHORegion,
	}

	res := new(domSchema.AddCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}

// Update Country
func (s *CountrySvc) Update(req *domSchema.UpdateCountryRequest, i identity.Identity) (*domSchema.UpdateCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryEtt domEntity.CountryEntity
	if err := dbCon.Where("code = ?", req.Keys.Code).First(&countryEtt).Error; err != nil || countryEtt.Code == "" {
		return nil, &sysErr.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Data not found (for Code=%s)", req.Keys.Code)}
	}

	// update
	countryEtt.Name = req.Data.Name
	countryEtt.ISO2Code = req.Data.ISO2Code
	countryEtt.ISO3Code = req.Data.ISO3Code
	countryEtt.WHORegion = req.Data.WHORegion

	countryEtt.UpdatedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Save(&countryEtt).Error; err != nil {
		if strings.Index(err.Error(), "Error 1062: Duplicate entry") > -1 {
			return nil, &sysErr.SystemError{StatusCode: http.StatusConflict, Err: err}
		}
		return nil, err
	}

	// response
	country := &domSchema.Country{
		ID:        countryEtt.ID,
		Code:      countryEtt.Code,
		Name:      countryEtt.Name,
		ISO2Code:  countryEtt.ISO2Code,
		ISO3Code:  countryEtt.ISO3Code,
		WHORegion: countryEtt.WHORegion,
	}

	res := new(domSchema.UpdateCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}

// Delete delete Country
func (s *CountrySvc) Delete(req *domSchema.DeleteCountryRequest, i identity.Identity) (*domSchema.DeleteCountryResponse, error) {
	// select db
	dbCon, err := s.handler.GetGormDB(s.dbConnectionName)
	if err != nil {
		return nil, err
	}

	// find
	var countryEtt domEntity.CountryEntity
	if err := dbCon.Where("code = ?", req.Code).First(&countryEtt).Error; err != nil || countryEtt.Code == "" {
		return nil, &sysErr.SystemError{StatusCode: http.StatusNotFound, Err: fmt.Errorf("Data not found (for Code=%s)", req.Code)}
	}

	// update deleted by user
	countryEtt.DeletedBy = fmt.Sprintf("%s@%s", i.Claims.Username, i.ClientDevices.IPAddress)

	if err := dbCon.Save(&countryEtt).Error; err != nil {
		return nil, err
	}

	// delete data
	if err := dbCon.Delete(&countryEtt).Error; err != nil {
		return nil, err
	}

	// response
	country := &domSchema.Country{
		ID:        countryEtt.ID,
		Code:      countryEtt.Code,
		Name:      countryEtt.Name,
		ISO2Code:  countryEtt.ISO2Code,
		ISO3Code:  countryEtt.ISO3Code,
		WHORegion: countryEtt.WHORegion,
	}

	res := new(domSchema.DeleteCountryResponse)
	res.Query = req
	res.Data = country

	return res, nil
}
