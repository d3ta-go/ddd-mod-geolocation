package application

import (
	appSvc "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/application/service"
	"github.com/d3ta-go/system/system/handler"
)

// NewGeoLocationApp new GeoLocationApp
func NewGeoLocationApp(h *handler.Handler) (*GeoLocationApp, error) {
	var err error
	app := new(GeoLocationApp)
	app.handler = h

	if app.CountrySvc, err = appSvc.NewCountrySvc(h); err != nil {
		return nil, err
	}

	return app, nil
}

// GeoLocationApp represent GeoLocationApp
type GeoLocationApp struct {
	handler    *handler.Handler
	CountrySvc *appSvc.CountrySvc
}
