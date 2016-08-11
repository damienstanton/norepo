//************************************************************************//
// API "cellar": Application Resource Href Factories
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

import "fmt"

// BottleHref returns the resource href.
func BottleHref(bottleID interface{}) string {
	return fmt.Sprintf("/bottles/%v", bottleID)
}
