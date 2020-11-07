package country

import (
	"encoding/json"

	domSchema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
)

// UpdateCountryReqDTO represent UpdateCountryReqDTO
type UpdateCountryReqDTO struct {
	Keys *UpdateCountryKeysDTO `json:"keys"`
	Data *UpdateCountryDataDTO `json:"data"`
}

type UpdateCountryKeysDTO domSchema.UpdateCountryKeys
type UpdateCountryDataDTO domSchema.UpdateCountryData

// UpdateCountryResDTO represent UpdateCountryResDTO
type UpdateCountryResDTO struct {
	Query interface{}        `json:"query"`
	Data  *domSchema.Country `json:"data"`
}

// ToJSON covert to JSON
func (r *UpdateCountryResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
