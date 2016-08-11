//************************************************************************//
// API "cellar": Application Media Types
//
// Generated with goagen v0.0.1, command line:
// $ goagen
// --out=$(GOPATH)/src/github.com/damienstanton/norepo/apiskeleton
// --design=github.com/damienstanton/norepo/apiskeleton/design
// --pkg=app
//
// The content of this file is auto-generated, DO NOT MODIFY
//************************************************************************//

package app

import "github.com/goadesign/goa"

// GoaExampleBottle media type.
//
// Identifier: application/vnd.goa.example.bottle+json
type GoaExampleBottle struct {
	// API href for making requests on the bottle
	Href string `json:"href" xml:"href"`
	// Unique bottle ID
	ID int `json:"id" xml:"id"`
	// Name of wine
	Name string `json:"name" xml:"name"`
}

// Validate validates the GoaExampleBottle media type instance.
func (mt *GoaExampleBottle) Validate() (err error) {
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}

	return err
}
