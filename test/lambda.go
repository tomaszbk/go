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
	fixtures := filepath.Join(runtime.GOROOT(), "test", "lambda.dir")
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
	dir, err := os.MkdirTemp("", "gon-lambda-exports-")
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
	write("main.go", []byte(`package main
import "featuretest/lib"
func main(){if lib.Hidden((v) => v.Value)!=13{panic("unnameable parameter type")}}
`))
	runAt(goTool, false, dir, "run", ".")

}
func checkInvalid(goTool string) {
	dir, err := os.MkdirTemp("", "gon-lambda-invalid-")
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
	{"no_target", `package main
var f=(x)=>x
func main(){}`},
	{"inferred_var", `package main
func f(){g:=(x)=>x;_ =g}
func main(){}`},
	{"no_function_target", `package main
var f any=(x)=>x
func main(){}`},
	{"parentheses", `package main
var f func(int)int=((x)=>x)
func main(){}`},
	{"parameter_count", `package main
var f func(int,int)int=(x)=>x
func main(){}`},
	{"duplicate", `package main
var f func(int,int)int=(x,x)=>x
func main(){}`},
	{"written_type", `package main
var f func(int)int=(x int)=>x
func main(){}`},
	{"bare_name", `package main
var f func(int)int=x=>x
func main(){}`},
	{"missing_arrow", `package main
var f func(int,int)int=(x,y)
func main(){}`},
	{"non_name", `package main
var f func(int)int=(1)=>1
func main(){}`},
	{"no_result_discard", `package main
var f func()=()=>1
func main(){}`},
	{"result_mismatch", `package main
var f func()int=()=>"x"
func main(){}`},
	{"result_count", `package main
func pair()(int,int){return 1,2};var f func()int=()=>pair()
func main(){}`},
	{"block_missing_return", `package main
var f func()int=()=>{}
func main(){}`},
	{"bare_return", `package main
var f func()int=()=>{return}
func main(){}`},
	{"propagation_boundary", `package main
func g()(int,error){return 0,nil};var f func()int=()=>g()!
func main(){}`},
	{"break_boundary", `package main
func f(){for {var g func()=()=>{break};_=g}}
func main(){}`},
	{"goto_boundary", `package main
func f(){L:var g func()=()=>{goto L};_=g}
func main(){}`},
	{"uninferred_parameters", `package main
func Apply[T any](f func(T)T)T{var z T;return z};var f=Apply((x)=>x+1)
func main(){}`},
	{"block_results_inference", `package main
func Map[T,U any](x []T,f func(T)U)[]U{return nil};var f=Map([]int{1},(x)=>{return x})
func main(){}`},
	{"untyped_nil_inference", `package main
func Pick[T any](f func()T)T{return f()};var f=Pick(()=>nil)
func main(){}`},
	{"direct_invocation", `package main
var x=(()=>1)()
func main(){}`},
	{"operator_context", `package main
var x=1+(()=>1)
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
