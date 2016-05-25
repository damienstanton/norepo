package test

import (
	"bytes"
	"fmt"
	"github.com/damienstanton/norepo/apiskeleton/app"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

// ShowBottleOK test setup
func ShowBottleOK(t *testing.T, ctrl app.BottleController, bottleID int) *app.GoaExampleBottle {
	return ShowBottleOKCtx(t, context.Background(), ctrl, bottleID)
}

// ShowBottleOKCtx test setup
func ShowBottleOKCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, bottleID int) *app.GoaExampleBottle {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/bottles/%v", bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	a, ok := resp.(*app.GoaExampleBottle)
	if !ok {
		t.Errorf("invalid response media: got %+v, expected instance of app.GoaExampleBottle", resp)
	}

	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}

	err = a.Validate()
	if err != nil {
		t.Errorf("invalid response payload: got %v", err)
	}
	return a

}

// ShowBottleNotFound test setup
func ShowBottleNotFound(t *testing.T, ctrl app.BottleController, bottleID int) {
	ShowBottleNotFoundCtx(t, context.Background(), ctrl, bottleID)
}

// ShowBottleNotFoundCtx test setup
func ShowBottleNotFoundCtx(t *testing.T, ctx context.Context, ctrl app.BottleController, bottleID int) {
	var logBuf bytes.Buffer
	var resp interface{}
	respSetter := func(r interface{}) { resp = r }
	service := goatest.Service(&logBuf, respSetter)
	rw := httptest.NewRecorder()
	req, err := http.NewRequest("GET", fmt.Sprintf("/bottles/%v", bottleID), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "BottleTest"), rw, req, nil)
	showCtx, err := app.NewShowBottleContext(goaCtx, service)
	if err != nil {
		panic("invalid test data " + err.Error()) // bug
	}
	err = ctrl.Show(showCtx)
	if err != nil {
		t.Fatalf("controller returned %s, logs:\n%s", err, logBuf.String())
	}

	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

}
