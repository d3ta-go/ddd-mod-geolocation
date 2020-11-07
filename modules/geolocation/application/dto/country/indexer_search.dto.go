package country

import (
	"encoding/json"

	domSchema "github.com/d3ta-go/ddd-mod-geolocation/modules/geolocation/domain/schema/country"
)

// SearchCountryIndexerReqDTO represent SearchCountryIndexerReqDTO
type SearchCountryIndexerReqDTO struct {
	domSchema.SearchCountryIndexerRequest
}

// SearchCountryIndexerResDTO represent SearchCountryIndexerResDTO
type SearchCountryIndexerResDTO struct {
	domSchema.SearchCountryIndexerResponse
}

// ToJSON covert to JSON
func (r *SearchCountryIndexerResDTO) ToJSON() []byte {
	json, err := json.Marshal(r)
	if err != nil {
		return nil
	}
	return json
}
