package kea

import "testing"

func TestStaticRoutes(t *testing.T) {
	if got := joinStaticRoutes([]string{"10.10.10.10,10.8.0.1", "10.10.10.11,10.8.50.1"}); got != "10.10.10.10,10.8.0.1,10.10.10.11,10.8.50.1" {
		t.Fatalf("join = %q", got)
	}
	for _, in := range []string{"10.10.10.10,10.8.0.1,10.10.10.11,10.8.50.1", "10.10.10.10,10.8.0.1;10.10.10.11,10.8.50.1", " 10.10.10.10 , 10.8.0.1 , 10.10.10.11 , 10.8.50.1 "} {
		r := parseStaticRoutes(in)
		if len(r) != 2 || r[0].DestinationIp.ValueString() != "10.10.10.10" || r[0].RouterIp.ValueString() != "10.8.0.1" || r[1].DestinationIp.ValueString() != "10.10.10.11" || r[1].RouterIp.ValueString() != "10.8.50.1" {
			t.Fatalf("parse(%q) = %v", in, r)
		}
	}
	if r := parseStaticRoutes(""); len(r) != 0 {
		t.Fatalf("parse(\"\") = %v", r)
	}
}
