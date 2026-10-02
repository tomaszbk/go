// run

//go:build !js && !wasip1 && gc

// Execute equivalent legacy and Gon programs, including a required unmodified
// Go baseline, and check invalid contexts and cross-package export bodies.
package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	goTool := filepath.Join(runtime.GOROOT(), "bin", "go")
	fixtures := filepath.Join(runtime.GOROOT(), "test", "nullsafety.dir")
	common := filepath.Join(fixtures, "common.go")
	legacyFile := filepath.Join(fixtures, "legacy.go")
	modernFile := filepath.Join(fixtures, "modern.go")
	baseline := os.Getenv("GON_BASELINE_GO")
	if baseline == "" {
		panic("GON_BASELINE_GO must name an unmodified compatible Go toolchain")
	}
	legacy := run(goTool, false, "run", common, legacyFile)
	compare("legacy/modern", legacy, run(goTool, false, "run", common, modernFile))
	compare("legacy/modern without inlining", legacy, run(goTool, false, "run", "-gcflags=-l", common, modernFile))
	compare("legacy baseline", legacy, run(baseline, true, "run", common, legacyFile))
	run(goTool, false, "vet", common, legacyFile)
	run(goTool, false, "vet", common, modernFile)
	if raceSupported() {
		compare("modern race detector", legacy, run(goTool, false, "run", "-race", common, modernFile))
	}
	checkExports(goTool, baseline, fixtures)
	checkInvalid(goTool)
}
func compare(what string, want, got []byte) {
	if !bytes.Equal(want, got) {
		panic(fmt.Sprintf("%s behavior differs\nwant:\n%s\ngot:\n%s", what, want, got))
	}
}

func run(goTool string, baseline bool, args ...string) []byte {
	return runAt(goTool, baseline, "", args...)
}

func runAt(goTool string, baseline bool, dir string, args ...string) []byte {
	cmd := exec.Command(goTool, args...)
	cmd.Dir = dir
	cmd.Env = append(os.Environ(), "GOENV=off", "GOTOOLCHAIN=local")
	if dir != "" {
		cmd.Env = append(cmd.Env, "GO111MODULE=on")
	}
	if baseline {
		var env []string
		for _, entry := range cmd.Env {
			if !strings.HasPrefix(entry, "GOROOT=") && !strings.HasPrefix(entry, "GOTOOLDIR=") {
				env = append(env, entry)
			}
		}
		cmd.Env = append(env, "GOFLAGS=")
	}
	out, err := cmd.CombinedOutput()
	if err != nil {
		panic(fmt.Sprintf("%s %v: %v\n%s", goTool, args, err, out))
	}
	return out
}

func checkExports(goTool, baseline, fixtures string) {
	dir, err := os.MkdirTemp("", "gon-nullsafety-exports-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	if err := os.Mkdir(filepath.Join(dir, "lib"), 0700); err != nil {
		panic(err)
	}
	write := func(name string, data []byte) {
		if err := os.WriteFile(filepath.Join(dir, name), data, 0600); err != nil {
			panic(err)
		}
	}
	copy := func(source, dest string) {
		data, err := os.ReadFile(filepath.Join(fixtures, source))
		if err != nil {
			panic(err)
		}
		write(dest, data)
	}
	write("go.mod", []byte("module featuretest\n\ngo 1.26\n"))
	copy("export_main.go", "main.go")
	copy("export_legacy.go", filepath.Join("lib", "lib.go"))
	legacy := runAt(goTool, false, dir, "run", ".")
	compare("exports baseline", legacy, runAt(baseline, true, dir, "run", "."))
	copy("export_modern.go", filepath.Join("lib", "lib.go"))
	compare("exports modern", legacy, runAt(goTool, false, dir, "run", "."))
	compare("exports modern without inlining", legacy, runAt(goTool, false, dir, "run", "-gcflags=-l", "."))
	output := runAt(goTool, false, dir, "build", "-gcflags=-m", "-o", filepath.Join(dir, "main.exe"), ".")
	if !strings.Contains(string(output), "inlining call to lib.Pick") {
		panic(fmt.Sprintf("cross-package inlining missing:\n%s", output))
	}
}
func checkInvalid(goTool string) {
	dir, err := os.MkdirTemp("", "gon-nullsafety-invalid-")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	for _, test := range invalidPrograms {
		file := filepath.Join(dir, test.name+".go")
		if err := os.WriteFile(file, []byte(test.source), 0600); err != nil {
			panic(err)
		}
		cmd := exec.Command(goTool, "build", "-o", filepath.Join(dir, "invalid.exe"), file)
		cmd.Env = append(os.Environ(), "GOENV=off", "GOTOOLCHAIN=local")
		out, err := cmd.CombinedOutput()
		if err == nil {
			panic(test.name + ": invalid program compiled")
		}
		for _, crash := range []string{"internal compiler error", "panic:", "unexpected type", "goroutine "} {
			if strings.Contains(string(out), crash) {
				panic(fmt.Sprintf("%s crashed compiler:\n%s", test.name, out))
			}
		}
	}
}

var invalidPrograms = []struct{ name, source string }{
	{"integer_guard", `package main
var x=1?.Field
func main(){}`},
	{"struct_guard", `package main
var x=struct{P *int}{}?.P
func main(){}`},
	{"nil_guard", `package main
var x=nil?.Field
func main(){}`},
	{"function_guard_type", `package main
var x=1?(2)
func main(){}`},
	{"assertion_guard", `package main
var x=any(nil)?.(int)
func main(){}`},
	{"non_nil_chain_value", `package main
var p *struct{N int};var x=p?.N
func main(){}`},
	{"non_nil_chain_bool", `package main
var p *struct{B bool};func f(){if p?.B{}}
func main(){}`},
	{"chain_assignment", `package main
var p *struct{P *int};func f(){p?.P=nil}
func main(){}`},
	{"chain_increment", `package main
var p *struct{N int};func f(){p?.N++}
func main(){}`},
	{"chain_address", `package main
var p *struct{P *int};var x=&p?.P
func main(){}`},
	{"chain_map_store", `package main
var p *struct{M map[int]int};func f(){p?.M[0]=1}
func main(){}`},
	{"defer_chain", `package main
var f func();func g(){defer f?()}
func main(){}`},
	{"go_chain", `package main
var f func();func g(){go f?()}
func main(){}`},
	{"guard_type_switch", `package main
var p *struct{V any};func f(){switch p?.V.(type){}}
func main(){}`},
	{"chain_comma_ok", `package main
var p *struct{M map[int]*int};var x,ok=p?.M[0]
func main(){}`},
	{"untyped_nil_coalesce", `package main
var x=nil ?? (*int)(nil)
func main(){}`},
	{"integer_coalesce", `package main
var x=0 ?? 1
func main(){}`},
	{"mismatched_coalesce", `package main
var p *int;var x=p ?? "x"
func main(){}`},
	{"implicit_dereference", `package main
var p *int;var x=p ?? 0
func main(){}`},
	{"assertion_coalesce", `package main
var a any;var x=a.(*int) ?? (*int)(nil)
func main(){}`},
	{"mixed_binary", `package main
var p *int;var x=*p ?? 0+1
func main(){}`},
	{"coalesce_constant", `package main
const x=(*int)(nil) ?? (*int)(nil)
func main(){}`},
	{"non_nil_assignment", `package main
func f(){x:=0;x ??=1}
func main(){}`},
	{"blank_assignment", `package main
func f(){_ ??=(*int)(nil)}
func main(){}`},
	{"chain_coalesce_assignment", `package main
var p *struct{P *int};func f(){p?.P ??=nil}
func main(){}`},
	{"multi_assignment", `package main
var p,q *int;func f(){p,q ??=nil}
func main(){}`},
	{"lone_question", `package main
var p *int;var x=p?
func main(){}`},
	{"guard_index", `package main
var p []*int;var x=p?[0]
func main(){}`},
}

// Match the supported Gon architectures for which Go supplies the race runtime.
func raceSupported() bool {
	switch runtime.GOARCH {
	case "amd64":
		return runtime.GOOS == "linux" || runtime.GOOS == "darwin" || runtime.GOOS == "windows" || runtime.GOOS == "freebsd" || runtime.GOOS == "netbsd"
	case "arm64":
		return runtime.GOOS == "linux" || runtime.GOOS == "darwin"
	}
	return false
}
