//************************************************************************//
// User Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/damienstanton/norepo/apiskeleton
// --design=github.com/damienstanton/norepo/apiskeleton/design
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package client

import (
	"github.com/goadesign/goa"
	"io"
)

// A bottle of wine
type GoaExampleBottle struct {
	// API href for making requests on the bottle
	Href string `json:"href" xml:"href"`
	// Unique bottle ID
	ID int `json:"id" xml:"id"`
	// Name of wine
	Name string `json:"name" xml:"name"`
}

// DecodeGoaExampleBottle decodes the GoaExampleBottle instance encoded in r.
func DecodeGoaExampleBottle(r io.Reader, decoderFn goa.DecoderFunc) (*GoaExampleBottle, error) {
	var decoded GoaExampleBottle
	err := decoderFn(r).Decode(&decoded)
	return &decoded, err
}
