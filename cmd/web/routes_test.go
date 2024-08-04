package main

import (
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/zaouldyeck/bookings/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		// Do nothing. No type errors encountered.
	default:
		t.Errorf("type is not *chi.Mux, type is %T", v)
	}
}
