package controller

import (
	"context"
	"net/http"
)

const key = "resWriter"

func setResponseWriter(ctx context.Context, w http.ResponseWriter) context.Context {
	res := context.WithValue(ctx, key, w)
	return res
}
