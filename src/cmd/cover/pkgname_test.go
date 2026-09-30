package main

import "testing"

func TestPackageName(t *testing.T) {
	var tests = []struct {
		fileName, pkgName string
	}{
		{"", ""},
		{"///", ""},
		{"fmt", ""}, // No Go file, improper form.
		{"fmt/foo.go", "fmt"},
		{"encoding/binary/foo.go", "binary"},
		{"encoding/binary/////foo.go", "binary"},
	}
	var tf templateFile
	for _, test := range tests {
		tf.Name = test.fileName
		td := templateData{
			Files: []*templateFile{&tf},
		}
		got := td.PackageName()
		if got != test.pkgName {
			t.Errorf("%s: got %s want %s", test.fileName, got, test.pkgName)
		}
	}
}
