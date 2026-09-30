package main

import "testplugin/issue44956/base"

func F() *map[int]int { return base.X }

func main() {}
