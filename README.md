# Gon

Gon is a programming language built by extending the Go compiler and toolchain
directly. It reduces repetitive code while retaining interoperability with Go
packages and the standard library.

The first implemented extension is error handling on conventional Go returns:

```go
func loadConfig(path string) (Config, error) {
    data := os.ReadFile(path)!
    return parseConfig(data)
}
```

Postfix `!` propagates errors; `or err { ... }` handles them locally. See the
[error-handling specification](design/error-handling/README.md) for semantics,
restrictions, executable examples, and validation instructions.

## Source files and compatibility

Gon source files use **`.go`**, including `_test.go` for tests. Modules retain
`go.mod` and `go.sum`. See the [language identity decision](design/language-identity.md).

Existing Go code remains supported, with one accepted syntax exception:
prefix `!` immediately followed by a newline is no longer supported. Same-line
negation continues to work. Code using Gon extensions requires this toolchain;
the upstream Go compiler does not understand the new syntax.

## Build and use

Build from source with a compatible Go bootstrap toolchain installed:

```sh
cd src
./make.bash
cd ..
./bin/go version
python3 misc/gon/build.py
python3 misc/gon/install.py
gon build ./path/to/package
gon fmt ./path/to/package
```

Set `GOROOT_BOOTSTRAP` if the build cannot locate the bootstrap toolchain.
The public commands are `gon` and `gonpls`. Private binaries retain their Go
names (`bin/go` and `bin/gofmt`) for internal tool compatibility.
Official downloads from go.dev are upstream Go, not Gon releases.

The fork includes syntax-aware parsing, type checking, formatting, and an initial
Gon language server based on gopls. This workspace configures the dedicated Gon
VS Code extension to use Gon without replacing tools in other projects. See
[tooling setup, tests, and current analysis limits](misc/gon/README.md).
For agents, scripts and CI, `gon query`, `gon check`, `gon refactor rename` and
`gon explain` expose the same semantic engine on the command line; see
[the command contract](misc/gon/CLI.md) and the project skill in
[`.agents/skills/gon`](.agents/skills/gon/SKILL.md).

## Upstream

Gon is derived from [Go](https://go.dev/). Upstream copyright notices and
the [license](LICENSE) are preserved.
