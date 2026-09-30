package main

var dec64Ops = []opData{}

var dec64Blocks = []blockData{}

func init() {
	archs = append(archs, arch{
		name:    "dec64",
		ops:     dec64Ops,
		blocks:  dec64Blocks,
		generic: true,
	})
}
