package service

import (
	"encoding/json"
	"fmt"
	"testing"

	appDTO "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/application/dto/country"
	"github.com/d3ta-go/system/system/config"
	"github.com/d3ta-go/system/system/handler"
	"github.com/d3ta-go/system/system/identity"
	"github.com/d3ta-go/system/system/initialize"
	"github.com/spf13/viper"
)

func newConfig(t *testing.T) (*config.Config, *viper.Viper, error) {
	c, v, err := config.NewConfig("../../../../conf")
	if err != nil {
		return nil, nil, err
	}
	if !c.CanRunTest() {
		panic(fmt.Sprintf("Cannot Run Test on env `%s`, allowed: %v", c.Environment.Stage, c.Environment.RunTestEnvironment))
	}
	return c, v, nil
}

func newCountrySvc(t *testing.T) (*CountrySvc, *handler.Handler, error) {
	h, err := handler.NewHandler()
	if err != nil {
		return nil, nil, err
	}

	c, v, err := newConfig(t)
	if err != nil {
		return nil, nil, err
	}

	h.SetDefaultConfig(c)
	h.SetViper("config", v)

	// viper for test-data
	viperTest := viper.New()
	viperTest.SetConfigType("yaml")
	viperTest.SetConfigName("test-data")
	viperTest.AddConfigPath("../../../../conf/data")
	viperTest.ReadInConfig()
	h.SetViper("test-data", viperTest)

	if err := initialize.LoadAllDatabaseConnection(h); err != nil {
		return nil, nil, err
	}
	if err := initialize.OpenAllIndexerConnection(h); err != nil {
		return nil, nil, err
	}

	r, err := NewCountrySvc(h)
	if err != nil {
		return nil, nil, err
	}

	return r, h, nil
}

func newIdentity(h *handler.Handler, t *testing.T) identity.Identity {
	i, err := identity.NewIdentity(
		identity.DefaultIdentity, identity.TokenJWT, "", nil, nil, h,
	)
	if err != nil {
		t.Errorf("NewIdentity: %s", err.Error())
	}
	if err := i.SetCasbinEnforcer("../../../../conf/casbin/casbin_rbac_rest_model.conf"); err != nil {
		t.Errorf("SetCasbinEnforcer: %s", err.Error())
	}

	i.Claims.Username = "test.d3tago"
	i.Claims.AuthorityID = "group:admin"

	return i
}

func TestCountrySvc_ListAll(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/countries/list-all"
	i.RequestInfo.RequestAction = "GET"

	resp, err := cSvc.ListAll(i)
	if err != nil {
		t.Errorf("ListAll: %s", err.Error())
	}

	if resp != nil {
		respJSON, err := json.Marshal(resp)
		if err != nil {
			t.Errorf("respJSON: %s", err.Error())
		}
		t.Logf("Resp: %s", respJSON)
	}
}

func TestCountrySvc_RefreshIndexer(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.app-layer.service.refresh-indexer")

	req := new(appDTO.RefreshCountryIndexerReqDTO)
	req.ProcessType = testData["processing-type"]

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/countries/indexer/refresh"
	i.RequestInfo.RequestAction = "POST"

	resp, err := cSvc.RefreshIndexer(req, i)
	if err != nil {
		t.Errorf("RefreshIndexer: %s", err.Error())
	}

	if resp != nil {
		t.Logf("Resp: %s", string(resp.ToJSON()))
	}
}

func TestCountrySvc_SearchIndexer(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.app-layer.service.search-indexer")

	req := new(appDTO.SearchCountryIndexerReqDTO)
	req.Name = testData["name"]

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/countries/indexer/search"
	i.RequestInfo.RequestAction = "POST"

	resp, err := cSvc.SearchIndexer(req, i)
	if err != nil {
		t.Errorf("SearchIndexer: %s", err.Error())
	}

	if resp != nil {
		t.Logf("Resp: %s", string(resp.ToJSON()))
	}
}

func TestCountrySvc_GetDetail(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.app-layer.service.get-detail")

	req := new(appDTO.GetCountryReqDTO)
	req.Code = testData["code"]

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/*"
	i.RequestInfo.RequestAction = "GET"

	resp, err := cSvc.GetDetail(req, i)
	if err != nil {
		t.Errorf("GetDetail: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestCountrySvc_Add(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.app-layer.service.add")

	req := new(appDTO.AddCountryReqDTO)
	req.Code = testData["code"]
	req.Name = testData["name"]
	req.ISO2Code = testData["iso2-code"]
	req.ISO3Code = testData["iso3-code"]
	req.WHORegion = testData["who-region"]

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country"
	i.RequestInfo.RequestAction = "POST"

	resp, err := cSvc.Add(req, i)
	if err != nil {
		t.Errorf("Add: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		// save to test-data
		// save result for next test
		viper.Set("test-data.geo-location.country.app-layer.service.update.id", resp.Data.ID)
		viper.Set("test-data.geo-location.country.app-layer.service.update.code", resp.Data.Code)
		viper.Set("test-data.geo-location.country.app-layer.service.delete.code", resp.Data.Code)
		if err := viper.WriteConfig(); err != nil {
			t.Errorf("Error: viper.WriteConfig(), %s", err.Error())
		}
		t.Logf("Resp: %s", string(respJSON))
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
	testData := viper.GetStringMapString("test-data.geo-location.country.app-layer.service.update")

	req := new(appDTO.UpdateCountryReqDTO)
	req.Keys = &appDTO.UpdateCountryKeysDTO{Code: testData["code"]}
	req.Data = &appDTO.UpdateCountryDataDTO{
		Name:      testData["name"],
		ISO2Code:  testData["iso2-code"],
		ISO3Code:  testData["iso3-code"],
		WHORegion: testData["who-region"],
	}

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/*"
	i.RequestInfo.RequestAction = "PUT"

	resp, err := cSvc.Update(req, i)
	if err != nil {
		t.Errorf("Update: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}

func TestCountrySvc_Delete(t *testing.T) {
	cSvc, h, err := newCountrySvc(t)
	if err != nil {
		t.Errorf("newCountrySvc: %s", err.Error())
	}

	viper, err := h.GetViper("test-data")
	if err != nil {
		t.Errorf("GetViper: %s", err.Error())
	}
	testData := viper.GetStringMapString("test-data.geo-location.country.app-layer.service.delete")

	req := new(appDTO.DeleteCountryReqDTO)
	req.Code = testData["code"]

	i := newIdentity(h, t)
	i.RequestInfo.RequestObject = "/api/v1/geolocation/country/*"
	i.RequestInfo.RequestAction = "DELETE"

	resp, err := cSvc.Delete(req, i)
	if err != nil {
		t.Errorf("Delete: %s", err.Error())
	}

	if resp != nil {
		respJSON := resp.ToJSON()
		t.Logf("Resp: %s", string(respJSON))
	}
}
