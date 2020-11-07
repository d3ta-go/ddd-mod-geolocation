package rdbms

import (
	"testing"

	schema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
	"github.com/d3ta-go/system/system/initialize"
)

func newCountrySvc(t *testing.T) (*CountrySvc, error) {
	h, err := newHandler(t)
	if err != nil {
		return nil, err
	}

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, err
	}

	r, err := NewCountrySvc(h)
	if err != nil {
		return nil, err
	}

	return r, nil
}

func TestCountrySvc_ListAll(t *testing.T) {
	cSvc, err := newCountrySvc(t)
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
	cSvc, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	req := &schema.GetDetailCountryRequest{Code: "XX"}
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
	cSvc, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	req := &schema.AddCountryRequest{
		Code:      "XX",
		Name:      "XX COUNTRY",
		ISO2Code:  "XX",
		ISO3Code:  "",
		WHORegion: "WPRO",
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
		t.Logf("Resp.AddCountryResponse: %s", string(respJSON))
	}
}

func TestCountrySvc_Update(t *testing.T) {
	cSvc, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	req := &schema.UpdateCountryRequest{
		Keys: &schema.UpdateCountryKeys{
			Code: "XX",
		},
		Data: &schema.UpdateCountryData{
			Name:      "XX COUNTRY UPDATED",
			ISO2Code:  "XX",
			ISO3Code:  "",
			WHORegion: "WPRO",
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
	cSvc, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
		return
	}

	req := &schema.DeleteCountryRequest{
		Code: "XX",
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
