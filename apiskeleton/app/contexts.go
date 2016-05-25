//************************************************************************//
// API "cellar": Application Contexts
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
	"strconv"
)

// ShowBottleContext provides the bottle show action context.
type ShowBottleContext struct {
	context.Context
	*goa.ResponseData
	*goa.RequestData
	Service  *goa.Service
	BottleID int
}

// NewShowBottleContext parses the incoming request URL and body, performs validations and creates the
// context used by the bottle controller show action.
func NewShowBottleContext(ctx context.Context, service *goa.Service) (*ShowBottleContext, error) {
	var err error
	req := goa.ContextRequest(ctx)
	rctx := ShowBottleContext{Context: ctx, ResponseData: goa.ContextResponse(ctx), RequestData: req, Service: service}
	paramBottleID := req.Params["bottleID"]
	if len(paramBottleID) > 0 {
		rawBottleID := paramBottleID[0]
		if bottleID, err2 := strconv.Atoi(rawBottleID); err2 == nil {
			rctx.BottleID = bottleID
		} else {
			err = goa.MergeErrors(err, goa.InvalidParamTypeError("bottleID", rawBottleID, "integer"))
		}
	}
	return &rctx, err
}

// OK sends a HTTP response with status code 200.
func (ctx *ShowBottleContext) OK(r *GoaExampleBottle) error {
	ctx.ResponseData.Header().Set("Content-Type", "application/vnd.goa.example.bottle")
	return ctx.Service.Send(ctx.Context, 200, r)
}

// NotFound sends a HTTP response with status code 404.
func (ctx *ShowBottleContext) NotFound() error {
	ctx.ResponseData.WriteHeader(404)
	return nil
}
