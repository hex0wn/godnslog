package server

import (
	"testing"
)

func TestParseDomain(t *testing.T) {
	var tests = []struct {
		Input         string
		Root          string
		ExpectPrefix  string
		ExpectRebind  bool
	}{
		{"712hu2c4gy34.godnslog.com", "godnslog.com", "712hu2c4gy34", false},
		{"www.godnslog.com", "godnslog.com", "www", false},
		{"r.www.godnslog.com", "godnslog.com", "r.www", false},
		{"r.godnslog.com", "godnslog.com", "r", true},
		{"a.r.godnslog.com", "godnslog.com", "a.r", true},
	}

	for i := 0; i < len(tests); i++ {
		test := &tests[i]
		prefix, rebind := parseDomain(test.Input, test.Root)
		if prefix != test.ExpectPrefix {
			t.Fatalf("test prefix(%v)!=expect(%v)", prefix, test.ExpectPrefix)
		}

		if rebind != test.ExpectRebind {
			t.Fatalf("test rebind(%v)!=ExpectRebind(%v)", rebind, test.ExpectRebind)
		}
	}
}
