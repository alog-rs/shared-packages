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

package utilities

import (
	"os"
)

type devEnv int

const (
	// EnvDev signifies a development environment
	EnvDev devEnv = iota
	// EnvProd signifies a production environment
	EnvProd = iota
)

// We set PRODUCTION_ENV when in production environments. We can assume
// that if this is not set we are in a development environment
func lookupEnv() devEnv {
	_, ok := os.LookupEnv("PRODUCTION_ENV")

	if ok {
		return EnvProd
	}

	return EnvDev
}

// IsDev returns if we are in a development environment
func IsDev() bool {
	return lookupEnv() == EnvDev
}

// IsProd returns if we are in a production environment
func IsProd() bool {
	return lookupEnv() == EnvProd
}

// Version returns the running version of the alog.rs service
func Version() string {
	result, ok := os.LookupEnv("VERSION")

	if !ok {
		return "unknown"
	}

	return result
}
