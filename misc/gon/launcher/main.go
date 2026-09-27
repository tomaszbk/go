// Copyright 2026 The Gon Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license in LICENSE.

// The public launchers keep the private Go-compatible tools off the user's PATH.
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	exe, err = filepath.EvalSymlinks(exe)
	if err != nil {
		return err
	}
	root := filepath.Clean(filepath.Join(filepath.Dir(exe), "..", ".."))
	name := strings.TrimSuffix(filepath.Base(exe), ".exe")
	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}
	program := filepath.Join(root, "bin", "go"+suffix)
	server := filepath.Join(root, "pkg", "tool", runtime.GOOS+"_"+runtime.GOARCH, "gonpls"+suffix)
	args := os.Args[1:]
	if name == "gonpls" {
		program = server
	} else if toolingCommand(args) {
		// Gon's additive tooling commands are implemented by the language
		// server's engine; every Go command keeps its upstream behavior.
		program, args = server, append([]string{"gon"}, args...)
	}
	if _, err := os.Stat(filepath.Join(root, "src", "go.mod")); err != nil {
		return fmt.Errorf("%s: cannot locate Gon toolchain relative to %s; keep gon/bin inside the toolchain or install a symlink to it: %w", name, exe, err)
	}
	if _, err := os.Stat(program); err != nil {
		return fmt.Errorf("%s: build the Gon tools first: %w", name, err)
	}
	// Pin both direct and subprocess invocations to Gon, even if the caller
	// has an upstream GOROOT or a downloaded toolchain configured globally.
	os.Setenv("GOROOT", root)
	os.Setenv("GOTOOLCHAIN", "local")
	os.Setenv("PATH", filepath.Join(root, "bin")+string(os.PathListSeparator)+os.Getenv("PATH"))
	return execProgram(program, args)
}

// toolingCommands are not Go commands, so forwarding them cannot change the
// behavior of any existing invocation.
var toolingCommands = map[string]bool{
	"query": true, "refactor": true, "check": true, "explain": true, "capabilities": true,
}

func toolingCommand(args []string) bool {
	if len(args) >= 2 && args[0] == "help" {
		return toolingCommands[args[1]] || args[1] == "tooling"
	}
	return len(args) > 0 && toolingCommands[args[0]]
}
