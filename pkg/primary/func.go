package primary

import (
	"encoding/json"
	"net/url"
)

// TableName имя таблицы для сущности Link
func (Link) TableName() string {
	return TableName
}

func (link Link) MarshalJSON() ([]byte, error) {
	URL := url.URL{
		Scheme:      link.Scheme,
		Opaque:      "",
		User:        nil,
		Host:        link.Host,
		Path:        link.Path,
		RawPath:     "",
		ForceQuery:  false,
		RawQuery:    link.Query,
		Fragment:    "",
		RawFragment: "",
	}

	return json.Marshal(&struct {
		FullURL string `json:"fullUrl"`
		UUID    string `json:"uuid"`
	}{
		UUID:    link.UUID,
		FullURL: URL.String(),
	})
}
