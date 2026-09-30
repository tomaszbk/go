// Package fipsdeps contains wrapper packages for internal APIs that are exposed
// to the FIPS module. Since modules are frozen upon validation and supported
// for a number of future versions, APIs exposed by crypto/internal/fips140deps/...
// must not be changed until the modules that use them are no longer supported.
package fipsdeps
