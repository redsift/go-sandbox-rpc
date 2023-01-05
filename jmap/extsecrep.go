package jmap

import "encoding/json"

type ExtSecReport struct {
	EHLO  json.RawMessage `json:"ehlo,omitempty"`
	TLS   json.RawMessage `json:"tls,omitempty"`
	SPF   json.RawMessage `json:"spf,omitempty"`
	DKIM  json.RawMessage `json:"dkim,omitempty"`
	DMARC json.RawMessage `json:"dmarc,omitempty"`
	ARC   json.RawMessage `json:"arc,omitempty"`
}
