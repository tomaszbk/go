package gonflow

import "strconv"

type node struct { value int }
func (n *node) Get(v int) *int { n.value = v; return &n.value }
func parse(s string) (int, error) { return strconv.Atoi(s) }
func load(s string) (*int, error) { v, err := parse(s); if err != nil { return nil, err }; return &v, nil }
func value(v int) *int { return &v }
