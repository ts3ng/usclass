package testgrp

import (
	"context"
	"math/rand"
	"net/http"

	"github.com/ardanlabs/service/foundation/web"
)

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100) % 2; n == 0 {
		//return errors.New("NON TRUSTED")
		//return v1Web.NewRequestError(errors.New("we trust this message"), http.StatusBadRequest)
		panic("test panic")
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
