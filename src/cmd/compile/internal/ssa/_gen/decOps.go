package main

var decOps = []opData{}

var decBlocks = []blockData{}

func init() {
	archs = append(archs, arch{
		name:    "dec",
		ops:     decOps,
		blocks:  decBlocks,
		generic: true,
	})
}
