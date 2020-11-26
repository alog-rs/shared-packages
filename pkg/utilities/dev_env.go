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
