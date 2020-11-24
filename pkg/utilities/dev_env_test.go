// Copyright (c) 2020, Chris Frank
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
//  * Redistributions of source code must retain the above copyright notice,
//    this list of conditions and the following disclaimer.
//  * Redistributions in binary form must reproduce the above copyright
//    notice, this list of conditions and the following disclaimer in the
//    documentation and/or other materials provided with the distribution.
//  * Neither the name of alog.rs nor the names of its contributors may be
//    used to endorse or promote products derived from this software without
//    specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

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

func TestVersion(t *testing.T) {
	mockVersion := "v0.0.1"
	cases := []struct {
		name           string
		versionPresent bool
		expected       string
	}{
		{"version present", true, mockVersion},
		{"version missing", false, "unknown"},
	}

	defer os.Unsetenv("VERSION")

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			os.Unsetenv("VERSION")

			if tc.versionPresent {
				os.Setenv("VERSION", mockVersion)
			}

			result := utilities.Version()

			if result != tc.expected {
				t.Errorf("got %s wanted %s", result, tc.expected)
			}
		})
	}
}
