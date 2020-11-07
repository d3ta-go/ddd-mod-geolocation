package model

// CountryEntity represent CountryEntity
type CountryEntity struct {
	ID        int64  `json:"ID" gorm:"primary_key;column:id"`
	Code      string `json:"code" gorm:"column:code;size:10;unique;not null"`
	Name      string `json:"name" gorm:"column:name;size:255;not null"`
	ISO2Code  string `json:"ISO2Code" gorm:"column:iso2code;size:5;not null"`
	ISO3Code  string `json:"ISO3Code" gorm:"column:iso3code;size:10"`
	WHORegion string `json:"WHORegion" gorm:"column:who_region;size:20"`

	BaseEntity
}

// TableName get real database table name
func (t *CountryEntity) TableName() string {
	return "mdm_geo_country"
}
