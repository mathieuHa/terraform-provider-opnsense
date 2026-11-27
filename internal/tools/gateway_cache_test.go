package tools

import (
	"errors"
	"testing"
)

func TestStaleGatewayError(t *testing.T) {
	for msg, want := range map[string]bool{
		"errors: map[route.gateway:Specify a valid gateway from the list matching the networks ip protocol.]": true,
		"errors: map[group.item:Option [tf_gw] not in list.]":                                                 true,
		"errors: map[route.network:Invalid network]":                                                          false,
	} {
		if got := staleGatewayError(errors.New(msg)); got != want {
			t.Fatalf("staleGatewayError(%q) = %v, want %v", msg, got, want)
		}
	}
}
