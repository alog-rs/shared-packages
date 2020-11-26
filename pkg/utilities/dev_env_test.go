package utilities_test

import (
	"os"
	"strconv"
	"testing"

	"github.com/alog-rs/shared-packages/pkg/utilities"
)

func TestIsDev(t *testing.T) {
	cases := []struct {
		name     string
		prod     bool
		expected bool
	}{
		{"in dev", false, true},
		{"in prod", true, false},
	}

	defer os.Unsetenv("PRODUCTION_ENV")

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv("PRODUCTION_ENV")

			if tc.prod {
				os.Setenv("PRODUCTION_ENV", "true")
			}

			result := utilities.IsDev()

			if result != tc.expected {
				t.Errorf("got %s wanted %s", strconv.FormatBool(result), strconv.FormatBool(tc.expected))
			}
		})
	}
}

func TestIsProd(t *testing.T) {
	cases := []struct {
		name     string
		prod     bool
		expected bool
	}{
		{"in dev", false, false},
		{"in prod", true, true},
	}

	defer os.Unsetenv("PRODUCTION_ENV")

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv("PRODUCTION_ENV")

			if tc.prod {
				os.Setenv("PRODUCTION_ENV", "true")
			}

			result := utilities.IsProd()

			if result != tc.expected {
				t.Errorf("got %s wanted %s", strconv.FormatBool(result), strconv.FormatBool(tc.expected))
			}
		})
	}
}
