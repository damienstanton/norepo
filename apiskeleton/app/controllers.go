//************************************************************************//
// API "cellar": Application Controllers
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

import (
	"github.com/goadesign/goa"
	"golang.org/x/net/context"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder(goa.NewJSONEncoder, "application/json")
	service.Encoder(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder(goa.NewXMLEncoder, "application/xml")
	service.Decoder(goa.NewJSONDecoder, "application/json")
	service.Decoder(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder(goa.NewJSONEncoder, "*/*")
	service.Decoder(goa.NewJSONDecoder, "*/*")
}

// BottleController is the controller interface for the Bottle actions.
type BottleController interface {
	goa.Muxer
	Show(*ShowBottleContext) error
}

// MountBottleController "mounts" a Bottle resource controller on the given service.
func MountBottleController(service *goa.Service, ctrl BottleController) {
	initService(service)
	var h goa.Handler

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		rctx, err := NewShowBottleContext(ctx, service)
		if err != nil {
			return err
		}
		return ctrl.Show(rctx)
	}
	service.Mux.Handle("GET", "/bottles/:bottleID", ctrl.MuxHandler("Show", h, nil))
	service.LogInfo("mount", "ctrl", "Bottle", "action", "Show", "route", "GET /bottles/:bottleID")
}
