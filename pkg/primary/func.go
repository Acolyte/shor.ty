package primary

import (
	"encoding/json"
	url2 "net/url"
)

// TableName имя таблицы для сущности Link
func (Link) TableName() string {
	return TableName
}

func (l Link) MarshalJSON() ([]byte, error) {
	url := url2.URL{
		Scheme:      l.Scheme,
		Opaque:      "",
		User:        nil,
		Host:        l.Host,
		Path:        l.Path,
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    l.Query,
		Fragment:    "",
		RawFragment: "",
	}

	return json.Marshal(&struct {
		FullURL string `json:"fullUrl"`
		UUID    string `json:"uuid"`
	}{
		UUID:    l.UUID,
		FullURL: url.String(),
	})
}
