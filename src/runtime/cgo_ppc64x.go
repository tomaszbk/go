//go:build ppc64 || ppc64le

package runtime

// crosscall1 calls into the runtime to set up the registers the
// Go runtime expects and so the symbol it calls needs to be exported
// for external linking to work.
//
//go:cgo_export_static _cgo_reginit
