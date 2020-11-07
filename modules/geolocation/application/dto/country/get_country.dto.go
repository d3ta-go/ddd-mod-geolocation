package country

import (
	"encoding/json"

	domSchema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
)

// GetCountryReqDTO represent GetCountryReqDTO
type GetCountryReqDTO struct {
	domSchema.GetDetailCountryRequest
}

// GetCountryResDTO represent GetCountryResDTO
type GetCountryResDTO struct {
	Query interface{}        `json:"query"`
	Data  *domSchema.Country `json:"data"`
}

// ToJSON covert to JSON
func (r *GetCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
