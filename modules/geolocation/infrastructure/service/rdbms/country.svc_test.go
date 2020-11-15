package rdbms

import (
	"testing"

	schema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/initialize"
)

func newCountrySvc(t *testing.T) (*CountrySvc, *handler.Handler, error) {
	h, err := newHandler(t)
	if err != nil {
		return nil, nil, err
	}

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}

	r, err := NewCountrySvc(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func TestCountrySvc_ListAll(t *testing.T) {
	cSvc, _, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	i := newIdentity(cSvc.handler, t)

	resp, err := cSvc.ListAll(i)
	if err != nil {
		t.Errorf("CountrySvc.ListAll: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp.ListAll: %s", string(respJSON))
	}
}

func TestCountrySvc_GetDetail(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.infra-layer.service.rdbms.get-detail")

	req := &schema.GetDetailCountryRequest{Code: testData["code"]}
	if err := req.Validate(); err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	i := newIdentity(cSvc.handler, t)

	resp, err := cSvc.GetDetail(req, i)
	if err != nil {
		t.Errorf("CountrySvc.GetDetail: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp.GetDetail: %s", string(respJSON))
	}
}

func TestCountrySvc_Add(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.infra-layer.service.rdbms.add")

	req := &schema.AddCountryRequest{
		Code:      testData["code"],
		Name:      testData["name"],
		ISO2Code:  testData["iso2-code"],
		ISO3Code:  testData["iso3-code"],
		WHORegion: testData["who-region"],
	}
	if err := req.Validate(); err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	i := newIdentity(cSvc.handler, t)

	resp, err := cSvc.Add(req, i)
	if err != nil {
		t.Errorf("CountrySvc.Add: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.infra-layer.service.rdbms.update.id", resp.Data.ID)
		viper.Set("test-data.geo-location.country.infra-layer.service.rdbms.update.code", resp.Data.Code)
		viper.Set("test-data.geo-location.country.infra-layer.service.rdbms.delete.code", resp.Data.Code)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp.AddCountryResponse: %s", string(respJSON))
	}
}

func TestCountrySvc_Update(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.infra-layer.service.rdbms.update")

	req := &schema.UpdateCountryRequest{
		Keys: &schema.UpdateCountryKeys{
			Code: testData["code"],
		},
		Data: &schema.UpdateCountryData{
			Name:      testData["name"],
			ISO2Code:  testData["iso2-code"],
			ISO3Code:  testData["iso3-code"],
			WHORegion: testData["who-region"],
		},
	}
	if err := req.Validate(); err != nil {
		t.Errorf("Validate: %s", err.Error())
		return
	}

	i := newIdentity(cSvc.handler, t)

	resp, err := cSvc.Update(req, i)
	if err != nil {
		t.Errorf("CountrySvc.Update: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp.UpdateCountryResponse: %s", string(respJSON))
	}
}

func TestCountrySvc_Delete(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.infra-layer.service.rdbms.delete")

	req := &schema.DeleteCountryRequest{
		Code: testData["code"],
	}
	if err := req.Validate(); err != nil {
		t.Errorf("Validate: %s", err.Error())
		return
	}

	i := newIdentity(cSvc.handler, t)

	resp, err := cSvc.Delete(req, i)
	if err != nil {
		t.Errorf("CountrySvc.Delete: %s", err.Error())
		return
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp.DelCountryResponse: %s", string(respJSON))
	}
}
