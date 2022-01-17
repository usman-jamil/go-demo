package controllers

import (
	"context"
	"encoding/json"
	"enttest/ent"
	"io"
	"net/http"
)

func RegisterControllers(ctx context.Context,
	client *ent.Client) {
	uc := newAircraftController(ctx, client)

	http.Handle("/aircrafts", *uc)
	http.Handle("/aircrafts/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {
	enc := json.NewEncoder(w)
	enc.Encode(data)
}
