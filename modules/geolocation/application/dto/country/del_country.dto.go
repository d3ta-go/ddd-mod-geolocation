package country

import (
	"encoding/json"

	domSchema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
)

// DeleteCountryReqDTO represent DeleteCountryReqDTO
type DeleteCountryReqDTO struct {
	domSchema.DeleteCountryRequest
}

// DeleteCountryResDTO represent DeleteCountryResDTO
type DeleteCountryResDTO struct {
	Query interface{}        `json:"query"`
	Data  *domSchema.Country `json:"data"`
}

// ToJSON covert to JSON
func (r *DeleteCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
