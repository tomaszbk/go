package main

var divmodOps = []opData{}

var divmodBlocks = []blockData{}

func init() {
	archs = append(archs, arch{
		name:    "divmod",
		ops:     divmodOps,
		blocks:  divmodBlocks,
		generic: true,
	})
}
