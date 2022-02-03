package testgrp

import (
	"context"
	"errors"
	"math/rand"
	"net/http"

	v1Web "github.com/ardanlabs/service/business/web/v1"
	"github.com/ardanlabs/service/foundation/web"
)

func Test(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	if n := rand.Intn(100) % 2; n == 0 {
		//return errors.New("NON TRUSTED")
		return v1Web.NewRequestError(errors.New("we trust this message"), http.StatusBadRequest)
	}

	status := struct {
		Status string
	}{
		Status: "OK",
	}
	return web.Respond(ctx, w, status, http.StatusOK)
}
