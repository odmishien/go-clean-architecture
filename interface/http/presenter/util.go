package presenter

import (
	"context"
	"net/http"
)

const key = "resWriter"

func getResponseWriter(ctx context.Context) http.ResponseWriter {
	rw, _ := ctx.Value("resWriter").(http.ResponseWriter)
	return rw
}
