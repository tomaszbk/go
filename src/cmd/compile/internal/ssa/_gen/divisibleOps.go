package main

var divisibleOps = []opData{}

var divisibleBlocks = []blockData{}

func init() {
	archs = append(archs, arch{
		name:    "divisible",
		ops:     divisibleOps,
		blocks:  divisibleBlocks,
		generic: true,
	})
}
