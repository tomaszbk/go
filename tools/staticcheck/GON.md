# Maintained Gon adaptation of honnef.co/go/tools

This is maintained source, not build output. Edit it directly. The original
module path, license and upstream tests are retained. UPSTREAM.json records the
imported upstream module and its module/go.mod checksums; those identify the
baseline, not the modified Gon tree.

Imported on 2026-10-02. x/tools uses the same baseline as cmd/vet, with the Gon
AST, CFG, SSA, type-code and constraint adaptations from the former tools and
cmd-vendor patches. Upstream now provides the V5 export reader that the old
patch supplied. Staticcheck preserves the previous IR changes, including
conditional expressions and the switch-tag fix. Future changes belong here,
not in patches or pkg/gon-tools.

The maintained modules use local replacements. Regenerate cmd vendor with
`python3 misc/gon/vendor.py` from the repository root and verify it with
`--check`. Run `GON_BASELINE_GO=/path/to/go python3 misc/gon/validate.py tooling`
for the integration gate. See misc/gon/INTEGRATION.md for feature obligations
and the distinction between structural traversal and semantic support.
